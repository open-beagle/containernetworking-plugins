// Copyright 2017 CNI authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"strings"

	"github.com/coreos/go-iptables/iptables"
	"github.com/mattn/go-shellwords"

	"github.com/containernetworking/plugins/pkg/utils"
)

type chain struct {
	table       string
	name        string
	entryChains []string // the chains to add the entry rule

	entryRules [][]string // the rules that "point" to this chain
	rules      [][]string // the rules this chain contains

	prependEntry bool // whether or not the entry rules should be prepended
}

// setup idempotently creates the chain. It will not error if the chain exists.
func (c *chain) setup(ipt *iptables.IPTables) error {
	err := utils.EnsureChain(ipt, c.table, c.name)
	if err != nil {
		return err
	}

	// Add the rules to the chain
	for _, rule := range c.rules {
		if err := utils.InsertUnique(ipt, c.table, c.name, false, rule); err != nil {
			return err
		}
	}

	// Add the entry rules to the entry chains
	for _, entryChain := range c.entryChains {
		for _, rule := range c.entryRules {
			r := []string{}
			r = append(r, rule...)
			r = append(r, "-j", c.name)
			if err := utils.InsertUnique(ipt, c.table, entryChain, c.prependEntry, r); err != nil {
				return err
			}
		}
	}

	return nil
}

// teardown idempotently deletes a chain. It will not error if the chain doesn't exist.
// It will first delete all references to this chain in the entryChains.
func (c *chain) teardown(ipt *iptables.IPTables) error {
	// flush the chain
	// This will succeed *and create the chain* if it does not exist.
	// If the chain doesn't exist, the next checks will fail.
	if err := utils.ClearChain(ipt, c.table, c.name); err != nil {
		return err
	}

	for _, entryChain := range c.entryChains {
		entryChainRules, err := ipt.List(c.table, entryChain)
		if err != nil || len(entryChainRules) < 1 {
			// Swallow error here - probably the chain doesn't exist.
			// If we miss something the deletion will fail
			continue
		}

		for _, entryChainRule := range entryChainRules[1:] {
			if strings.HasSuffix(entryChainRule, "-j "+c.name) {
				chainParts, err := shellwords.Parse(entryChainRule)
				if err != nil {
					return fmt.Errorf("error parsing iptables rule: %s: %v", entryChainRule, err)
				}
				chainParts = chainParts[2:] // List results always include an -A CHAINNAME

				if err := utils.DeleteRule(ipt, c.table, entryChain, chainParts...); err != nil {
					return err
				}

			}
		}
	}

	return utils.DeleteChain(ipt, c.table, c.name)
}

// check the chain.
func (c *chain) check(ipt *iptables.IPTables) error {
	exists, err := utils.ChainExists(ipt, c.table, c.name)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("chain %s not found in iptables table %s", c.name, c.table)
	}

	for i := len(c.rules) - 1; i >= 0; i-- {
		match := checkRule(ipt, c.table, c.name, c.rules[i])
		if !match {
			return fmt.Errorf("rule %s in chain %s not found in table %s", c.rules, c.name, c.table)
		}
	}

	for _, entryChain := range c.entryChains {
		for i := len(c.entryRules) - 1; i >= 0; i-- {
			r := []string{}
			r = append(r, c.entryRules[i]...)
			r = append(r, "-j", c.name)
			matchEntryChain := checkRule(ipt, c.table, entryChain, r)
			if !matchEntryChain {
				return fmt.Errorf("rule %s in chain %s not found in table %s", c.entryRules, entryChain, c.table)
			}
		}
	}

	return nil
}

func checkRule(ipt *iptables.IPTables, table, chain string, rule []string) bool {
	exists, err := ipt.Exists(table, chain, rule...)
	if err != nil {
		return false
	}
	return exists
}
