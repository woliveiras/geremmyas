# LaTeX Conventions for This Repository

Conventions established in `papers/2026-slm-android-litert/paper.tex`. Follow these for all new papers.

---

## File Structure

```
papers/YYYY-short-slug/
├── README.md           # Abstract, metadata table, build instructions
├── paper.tex           # Main manuscript
├── title-page.tex      # Standalone title page for double-blind packet
├── references.bib      # Bibliography
└── scripts/            # Analysis/benchmark scripts (MIT licensed)
```

---

## Document Class and Core Packages

```latex
\documentclass[11pt,a4paper]{article}

\usepackage[utf8]{inputenc}
\usepackage[T1]{fontenc}
\usepackage{lmodern}
\usepackage[english]{babel}
\usepackage{geometry}
\geometry{margin=1in}
\usepackage{graphicx}
\usepackage{booktabs}       % \toprule \midrule \bottomrule in tables
\usepackage{hyperref}
\usepackage{amsmath}
\usepackage{cleveref}       % \cref{} - use instead of \ref{} everywhere
\usepackage{enumitem}
\usepackage{xcolor}
\usepackage{listings}
\usepackage{authblk}
\usepackage{microtype}      % Always include - improves typography
\usepackage{tabularx}
\usepackage{longtable}      % For multi-page tables
```

---

## Anonymization Toggle

The anonymization system uses a single boolean flag. Change only one line to switch between arXiv/preprint and double-blind submission:

```latex
% --- Anonymization toggle ---
% Set \anonymizedtrue  for double-blind submission
% Set \anonymizedfalse for arXiv/preprint/camera-ready
\newif\ifanonymized
\anonymizedfalse   % <-- CHANGE THIS LINE ONLY
```

**Usage pattern throughout the document:**

```latex
% Author block
\ifanonymized
  \author{[Anonymous for review]}
  \date{}
\else
  \author{William Oliveira\\
  Independent Researcher\\
  \texttt{contact@woliveiras.com}}
  \date{April 2026}
\fi

% Self-citations / repository URLs
\ifanonymized
  the source repository~[anonymous]
\else
  the Palabrita repository~\cite{palabrita}
\fi

% Acknowledgments (entire section hidden)
\ifanonymized
  % [Acknowledgments hidden for blind review]
\else
  \section*{Acknowledgments}
  ...
\fi
```

**Important:** Also suppress PDF metadata in blind mode:
```latex
\hypersetup{
  colorlinks=true,
  linkcolor=blue!70!black,
  citecolor=blue!70!black,
  urlcolor=blue!70!black,
  \ifanonymized
    pdfauthor={},
    pdftitle={},
  \fi
}
```

---

## `title-page.tex` - Separate Title Page

Used for double-blind submission packets where the main `paper.tex` has `\anonymizedtrue`, but the submission system requires a separate deanonymized cover page.

```latex
\documentclass[11pt,a4paper]{article}
\usepackage[utf8]{inputenc}
\usepackage[T1]{fontenc}
\usepackage{lmodern}
\usepackage{geometry}
\geometry{margin=1in}
\usepackage{hyperref}

\begin{document}
\thispagestyle{empty}
\begin{center}
  \vspace*{2cm}
  {\LARGE\bfseries [Full Paper Title]\par}
  \vspace{2em}
  {\large William Oliveira\par}
  \vspace{0.5em}
  {\normalsize Independent Researcher\par}
  {\normalsize ORCID: 0000-0000-0000-0000\par}
  {\normalsize \href{mailto:contact@woliveiras.com}{contact@woliveiras.com}\par}
  \vspace{2em}
  {\normalsize \textbf{Funding:} None\par}
  {\normalsize \textbf{Conflicts of interest:} None\par}
  {\normalsize \textbf{Data availability:} [URL or statement]\par}
\end{center}
\end{document}
```

---

## Build Commands

```bash
# Standard build (no bibliography)
pdflatex paper.tex && pdflatex paper.tex

# Build with bibliography
pdflatex paper.tex
bibtex paper
pdflatex paper.tex
pdflatex paper.tex

# Or with biber (biblatex)
pdflatex paper.tex
biber paper
pdflatex paper.tex
pdflatex paper.tex

# Check for double-blind leaks (after \anonymizedtrue build)
pdfinfo paper.pdf | grep -E "Author|Title|Creator"
grep -n "\\\\cite{" paper.tex | grep -v "\\\\ifanonymized"  # rough check
```

---

## Tables

Use `booktabs` (no vertical lines; horizontal lines with `\toprule`, `\midrule`, `\bottomrule`):

```latex
\begin{table}[ht]
\centering
\small
\caption{Caption goes above the table.}
\label{tab:my-label}
\begin{tabularx}{\textwidth}{lXX}
\toprule
\textbf{Col 1} & \textbf{Col 2} & \textbf{Col 3} \\
\midrule
Row 1 & Value & Value \\
Row 2 & Value & Value \\
\bottomrule
\end{tabularx}
\end{table}
```

- `\small` inside table environment to reduce font size
- Caption ABOVE for tables, BELOW for figures
- Use `\label{tab:...}` prefix for tables, `\label{fig:...}` for figures, `\label{sec:...}` for sections

---

## Cross-References

Always use `\cref{}` (from `cleveref`) instead of bare `\ref{}`:

```latex
% Good
As shown in \cref{tab:framework-comparison}...
\cref{sec:background} reviews...

% Avoid
Table~\ref{tab:framework-comparison}
Section~\ref{sec:background}
```

---

## Listings (Code Blocks)

```latex
\lstset{
  basicstyle=\small\ttfamily,
  breaklines=true,
  frame=single,
  columns=fullflexible,
  aboveskip=8pt,
  belowskip=8pt,
}

% Usage
\begin{lstlisting}[language=Kotlin, caption={Description.}, label={lst:example}]
fun example(): String = "hello"
\end{lstlisting}
```

---

## Licensing

- **Paper text** (`.tex`, `.pdf`): CC BY 4.0
- **Scripts / code** (`.py`, `.sh`, `.kt`, etc.): MIT

Each paper folder requires:
- `LICENSE` - contains CC BY 4.0 text for the paper
- Root `LICENSE` - contains MIT for the repository (scripts)

---

## `.gitignore` Entries for LaTeX

The root `.gitignore` should exclude:
```
*.aux
*.log
*.out
*.toc
*.bbl
*.blg
*.fdb_latexmk
*.fls
*.synctex.gz
*.pdf        # Build outputs; source tracked, PDFs not
```

---

## Naming Convention

| Item | Convention | Example |
|---|---|---|
| Folder | `YYYY-short-slug` | `2026-slm-android-litert` |
| Short slug | hyphen-separated, lowercase, ≤4 words | `slm-arm-sbcs` |
| Section labels | `sec:keyword` | `\label{sec:threats}` |
| Table labels | `tab:keyword` | `\label{tab:framework-comparison}` |
| Figure labels | `fig:keyword` | `\label{fig:architecture}` |
| Listing labels | `lst:keyword` | `\label{lst:system-prompt}` |
| Equation labels | `eq:keyword` | `\label{eq:energy-per-inference}` |
