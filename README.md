# Komerce - Backend Coding Test Solutions

This repository contains the solutions for the **Komerce Backend Coding Test**, implemented in **Go**.

## 📋 Table of Contents
1. [Problem 1: Sort Characters](#1-problem-sort-characters)
2. [Problem 2: PSBB (Pembatasan Sosial Berskala Besar)](#2-problem-psbb)
3. [Prerequisites](#prerequisites)
4. [How to Run](#how-to-run)

---

## 💻 Problem Overview

### 1. Problem: Sort Characters
Processes an input string to separate vowels and consonants based on the following rules:
- Preserves the original order of appearance.
- Treats all characters as lowercase and ignores whitespaces.
- Duplicates characters based on their frequency of occurrence in the original string.

### 2. Problem: PSBB
Determines the minimum number of minibuses required to transport family members with strict capacity rules:
- Maximum capacity per minibus: 4 passengers.
- Members of the same family must ride in the same bus.
- One bus can carry members of at most **two different families**.
- Validates that the number of families (`n`) matches the provided family members' input.

---

## 🛠️ Prerequisites

Make sure you have **Go** installed on your system:
- Go 1.18 or higher.

---

## 🚀 How to Run

### Running Problem 1 (Sort Characters)
`go run sort-characters/main.go`

### Running Problem 2 (PSBB)
`go run psbb/main.go`

---
