# Data Presentation for Scientific Papers

Guidelines for figures, tables, and benchmark data visualization. Used during `data-presentation` mode.

---

## Core Principles

1. **Every figure must earn its place**: a figure is justified only if it conveys a pattern that is harder to see in a table or text.
2. **Show variance, not just central tendency**: a bar chart showing only means conceals variance and is misleading.
3. **Make comparisons easy**: arrange data so the comparison the reader needs to make is visually adjacent.
4. **Label everything**: axes, units, sample sizes, significance markers.
5. **Design for print and accessibility**: grayscale-readable; color-blind-safe palettes.

---

## Chart Type Selection Guide

| Data type | Recommended | Avoid |
|---|---|---|
| Distribution of continuous data | Box plot, violin plot | Bar chart |
| Comparison of means with variance | Box plot with jitter or violin | Bar chart with error bars |
| Trend over time / iterations | Line chart | Bar chart |
| Correlation between two continuous variables | Scatter plot + regression line | Heatmap |
| Multiple conditions × multiple metrics | Faceted plots (one panel per metric) | Single overcrowded chart |
| Ranking / league table | Horizontal bar chart (sorted) | Pie chart |
| Proportion of whole (few categories) | Stacked bar | Pie chart (>4 categories) |
| Distribution comparison across groups | Overlapping density plots / violin | Separate histograms |

### Why to avoid bar charts for distributions

A bar chart showing mean ± SD hides the distribution shape. Two datasets can have the same mean and SD but completely different distributions (the "datasaurus dozen" effect). Use:
- **Box plot**: shows median, IQR, whiskers, outliers
- **Violin plot**: shows full distribution shape + box plot overlay
- **Beeswarm / strip plot**: shows individual data points when N ≤ 50

---

## Benchmark Data Visualization

### tokens/s or throughput comparison across conditions

**Best**: Horizontal bar chart (sorted descending), one bar per condition, with error bars (±1 SD or 95% CI).

```
Condition A  ████████████████ 45.2 ± 3.1
Condition B  ████████████ 38.7 ± 2.8
Condition C  █████████ 29.3 ± 4.2
```

- Sort by performance (best at top)
- Show N in legend or caption ("N=30 runs per condition")
- Error bars: state whether SD or 95% CI

**Avoid**: Grouping too many conditions on one chart. If >8 conditions, split into multiple panels.

### Latency distribution

**Best**: Box plot per condition. Shows median, quartiles, outliers.
- If N ≥ 30: add violin layer behind box plot
- Mark outliers (IQR × 1.5 rule) as individual points

**Avoid**: Reporting only mean latency without variance - one outlier run distorts the mean significantly.

### Energy per inference

**Best**: Scatter plot (tokens/s on x-axis, energy/inference on y-axis) - reveals the efficiency frontier.
- Different marker shapes for different model families
- Different colors for different SBCs / hardware
- Add Pareto frontier line

### Accuracy × throughput trade-off

**Best**: Scatter plot with metric1 on x, metric2 on y. Each point = one model/condition.
- Label outliers (best/worst per axis)
- Quadrant lines at meaningful thresholds (e.g., "usable throughput" threshold)

### Temperature curve over time

**Best**: Line chart (time on x, temperature on y). One line per run or per condition.
- Mark cool-down threshold
- Mark measurement window

---

## Statistical Significance Markers

| Marker | Meaning | Use when |
|---|---|---|
| `*` | p < 0.05 | Significant at standard threshold |
| `**` | p < 0.01 | Highly significant |
| `***` | p < 0.001 | Very highly significant |
| `ns` | Not significant | Always include - absence of marker is ambiguous |
| Effect size | Cohen's d, r, η² | Always alongside p-value |

**Rule**: p-values without effect sizes are incomplete. A p < 0.001 result with Cohen's d = 0.1 is statistically significant but practically trivial.

---

## Tables

### When to use tables vs. figures

- **Table**: exact values matter; reader needs to look up specific numbers; comparing >3 conditions on >3 metrics
- **Figure**: pattern or trend is the point; relative magnitude is more important than exact value

### Table design rules

- Caption ABOVE the table (LaTeX convention; also ACM/IEEE style)
- Use `booktabs` package (no vertical lines; `\toprule`, `\midrule`, `\bottomrule`)
- Right-align numeric columns; align decimal points
- Bold the best value per row (helps reader scan)
- Include units in column header, not in every cell
- For large tables: use `\small` font size; consider rotating with `\rotatebox` or `rotating` package

**Example header:**
```
| Model        | tokens/s (mean ± SD) | Energy/inf (mJ) | Accuracy (%) |
```

Not:
```
| Model        | Speed          | Power        | Quality |
```

### Comparison table (related work)

Use a matrix where rows = related papers, columns = dimensions of comparison.
- Mark each cell: ✓ (covered), △ (partially covered), ✗ (not covered), ? (unclear)
- Add your paper as the last row to make the gap explicit

---

## Axis Labels and Units

Mandatory:
- Both axes labeled with quantity name AND unit: "Inference latency (ms)", "Energy per inference (mJ)", "Memory bandwidth (GB/s)"
- Axis range starts at 0 for absolute values (unless variation is the point - then explain)
- If axis range does not start at 0, add a note in the caption

Forbidden:
- Axes labeled "Value" or "Score" without specifying what value or what score
- Y-axis break (shown as `//`) without a very strong justification and explicit notation
- Missing units on axes with physical quantities

---

## Error Bars: Be Explicit

Every error bar must state in the caption or legend what it represents:
- Standard deviation (SD) - describes data spread
- Standard error of the mean (SEM) - describes uncertainty about the mean; DO NOT use to imply narrow variance
- 95% confidence interval (CI) - preferred for hypothesis testing context

**Anti-pattern**: Publishing bar charts with error bars labeled only "Error bars" or with no label at all.

---

## Color Accessibility

1. **Do not rely on color alone** to convey information - use shape, pattern, or direct label too.
2. **Color-blind-safe palettes** (good for up to 8 categories):
   - Okabe-Ito: `#E69F00, #56B4E9, #009E73, #F0E442, #0072B2, #D55E00, #CC79A7, #000000`
   - ColorBrewer qualitative palettes (colorbrewer2.org)
3. Test with a grayscale print preview - if the chart is unreadable in grayscale, fix it.
4. Avoid: red vs. green (deuteranopia); blue vs. purple (tritanopia)

---

## Figure Captions

A good caption stands alone - the reader can understand the figure without reading the paper body.

**Template:**
```
Figure N. [What is shown]. [Key finding or pattern]. [N = X observations / runs; conditions described].
[What error bars represent, if applicable]. [How to read the chart, if not obvious].
```

**Example:**
```
Figure 3. Inference throughput (tokens/s) for all model × quantization × SBC combinations,
sorted by mean throughput. Error bars show ±1 standard deviation (N=30 runs per condition).
The Radxa Rock 5B consistently achieves higher throughput than both Raspberry Pi 5 and
Orange Pi 5 for models with Q4_K_M quantization, attributable to its higher memory bandwidth (51 GB/s).
```

---

## Common Data Presentation Errors

| Error | Why it is a problem | Fix |
|---|---|---|
| Bar chart for distributions | Hides variance, distribution shape | Box plot or violin plot |
| Y-axis starts at non-zero without note | Makes differences look larger than they are | Start at 0 or add explicit break note |
| SEM bars instead of SD | SEM converges to 0 with large N, looks precise but misleads | Use SD or 95% CI; label explicitly |
| Missing units on axes | Reader cannot interpret the chart | Add units to all axes with physical quantities |
| Unlabeled error bars | Ambiguous: SD? SEM? CI? | Always label in caption or legend |
| No effect sizes with p-values | Statistical significance ≠ practical significance | Report Cohen's d or equivalent alongside p |
| Single color with no pattern distinction | Not accessible to color-blind readers | Add marker shapes or patterns |
| Over-plotted scatter with no jitter | Points overlap, true density invisible | Add jitter; or use hexbin for large N |
| Table with too many significant digits | 45.23847 tokens/s implies false precision | Round to appropriate precision (e.g., 45.2 ± 3.1) |
