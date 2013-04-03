# Inlämningsmall

Det här är en enklare version av Frosts `assignment_template` för att uppmuntra
användning av LaTeX för inlämningar i kursen DD1339: Introduktion till
datalogi.

## Användning

Kopiera `template.tex` till t.ex. `HT1.tex` (för första uppgiften). Ändra sedan
i den nya LaTeX-filen. `template.tex` innehåller bland annat
exempel på hur man kan inkludera kodsnuttar eller hela `.java`-filer i
inlämningen.

Kompilera sedan `.tex`-filen genom att köra `make HT1.pdf` (i detta exempel),
eller kompilera den direkt på vanligt sätt.

> *Kommer du på andra bra exempelsaker att ha med i denna mall så skriv en
> Issue på GitHub och föreslå. Vill du fixa till själv, så forka projektet och
> begär en pull-request sedan. :)*

## Installation

Enklast är att trycka på `Downloads` och hämta ner projektet som en zip-fil.
Annars kan man klona projektet med git:

> `git clone git://github.com/pbos/assignment_template_light.git`

Skapa jättegärna en egen fork av projektet (om du förstår vad det betyder) och
ladda ner den. Då kan du ändra i ditt projekt och sedan begära en pull-request
för att få in förändringar i denna mall.

Sedan är det bara att tuta och köra. Förutsatt att man har en
LaTeX-installation med de paket och som behövs för att köra.

## Inställningar

### LaTeX-miljö

För att ändra inställningar, kopiera filen `latex_env.example` till
`latex_env`. Där går det att ställa in om man ska köra t.ex. `xelatex`
eller `pdflatex`.

### Namn/Asse-inställningar

Kopiera `config.tex.example` till `config.tex` och byt ut namn/gruppnummer etc.
i filen till dina egna.

## LaTeX-paket

Saknas instruktioner för hur man gör i just ditt system, berätta hur man gör
eller forka och fixa. :)

### Ubuntu

På Ubuntu 11.08/12.04 räcker det att köra

> `sudo apt-get install texlive-latex-base texlive-latex-recommended`

### Windows

Verkar fungera med [TeXworks](http://www.tug.org/texworks/)
default-installation.

### Mac OS X

MacTex verkar vara standarddistrubitionen av TeXLive för Mac. Kan hittas
[här](http://www.tug.org/mactex/) och verkar fungera bra. MacTeX.pkg är den
intressanta filen.
