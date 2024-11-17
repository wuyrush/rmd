# `rmd`: Render Markdown Doc

Handy tool for me to render and preview Markdown docs.

NOTE: Preview feature supports only OSX for now. Meanwhile, OSX users will need to configure the application for file type `public.html` to web browser in advance (most times this has been done) for `rmd` to work. More on [SO](https://stackoverflow.com/questions/10006958/open-an-html-file-with-default-browser-using-bash-on-mac).

## Motivation

As an author who writes docs in Markdown syntax, I want a convenient, off-line way to render my Markdown documents to HTML and view it.

With `rmd`, I am able to:
1. Start rendering from my editor (e.g. a leader key shortcut in vim) and view result in my web browser.
2. Start rendering from my editor and get rendered text.
3. Start rendering from my editor and save rendered result to a new file.
4. Start rendering from command line and output result to stdout or a new file.
5. Start rendering from command line and view result in web browser.

Plus:
1. If I am only previewing my doc, I shouldn't need to bother with any clean-up of temporary files generated for preview;
2. For better visuals I can style rendered document with themes e.g. Github Markdown light theme.

## Usage

```
# Preview w/ style
rmd -preview -style -i <fp> 

# output w/ style
rmd -style -i <fp> > out.html

# output plain rendered data
rmd -i <fp> > out.html

# read from stdin and output to stdout
rmd
```
