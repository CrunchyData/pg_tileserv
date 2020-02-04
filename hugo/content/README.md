## Content Organization

The `content` folder is a collection of Markdown files placed in different subfolders (plus a top-level `_index.md` file). When the site renders, these subfolders will display as top-level sections accessible in the navigation.

`_index.md` is for the guide's homepage.

A section (subfolder) may also contain multiple subtopics. 
For example, the *Installation* section has three subtopics: *Requirements* (`05-requirements.md`), *Installation* (`06-installation.md`), and *Deployment* (`07-deployment.md`). On the other hand, the *Learn More* section does not have any subtopics, so it contains only one file (the index file). Each subfolder must have its own index file. 

## Adding Content

Each `.md` file will start with a snippet that looks like this:
```
---
title: "Page Title"
date: 
draft: false
weight: 10
---
```
`title` corresponds to the title of the page, or section if it is an index file within a subfolder. `date` is optional. `draft` indicates the draft status: set to `true` if you don't want Hugo to publish the page.

`weight` determines the section and page order. Each section is ordered based on their respective index files. Use double digits (e.g. `weight: 10`) to order the top level sections. The subtopics are ordered based on the individual Markdown files. Use more digits (e.g. `weight: 150`) for ordering subtopics.

You should feel free to remove or add new files or subfolders/sections as you see fit.

If you would like to add a new file (i.e. a separate page), make sure to follow the file naming convention: 

- prefixed with a digit to help you keep track of the order (e.g. `09-`)
- dashes instead of spaces
- use the `.md` extension

If necessary, change the number prefixes of the other files to apply a change in order.
