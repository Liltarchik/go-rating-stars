# Score-to-Stars Rating Generator (Go) ⭐

A lightweight Go CLI utility and helper function that converts numerical percentage scores ($0-100$) into a visual 5-star Unicode rating format.

## 🚀 Overview & Logic

The `Rating_stars` function divides input values into 5 uniform tiers ($20$ points per star):

| Score Range | Output | Rating Tier |
| :--- | :--- | :--- |
| `0 – 19` | `☆☆☆☆☆` | Tier 0 |
| `20 – 39` | `★☆☆☆☆` | Tier 1 |
| `40 – 59` | `★★☆☆☆` | Tier 2 |
| `60 – 79` | `★★★☆☆` | Tier 3 |
| `80 – 100` | `★★★★★` | Tier 5 (Tier 4 at 80–99: `★★★★☆`) |
| `< 0` or `> 100` | `☆☆☆☆☆` | Out of range / Default |

## 🛠️ Prerequisites & Installation

* Go 1.18+

1. **Clone the repository:**
   ```bash
   git clone [https://github.com/Liltarchik/go-rating-stars.git](https://github.com/Liltarchik/go-rating-stars.git)
   cd go-rating-stars
