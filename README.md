# Minimum Window Substring

> NeetCode 150 · Sliding Window · Blind 75 — expand until the window covers t, then contract while it still does; a required-distinct counter makes "covers" O(1). Go, stdlib only, tested.

## Problem

The shortest substring of s containing all of t's characters *with multiplicity* (`"a"` needs one a, `"aa"` needs two). `ADOBECODEBANC` / `ABC` → `BANC`. The hardest of the classic window problems.

## Approach

The variable window's two-phase rhythm: march the right edge until the window satisfies t, then pull the left edge in while satisfaction holds — every position of the left edge inside the loop yields a *locally minimal* window, one of which is the global answer. The O(1) satisfaction test is the `required` counter: how many distinct characters of t are still short. It changes only when a count crosses its exact threshold — `have == need` on entry (one fewer required), `have < need` on exit (one more) — so neither phase ever rescans counts. Both edges only move forward: 2n steps total.

**Complexity:** O(n + m) time, O(128) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
