# Template: User Guide  

This template is for Crunchy Data employees to use as a reference when creating a User Guide. The easiest way to use this is to download the repository and import the files/content into your own project. 

We will assume that you are developing docs locally and will also be testing in browser.

We use [Hugo](https://gohugo.io/getting-started/installing/) to generate docs pages.

> Use version [< 0.60](https://github.com/gohugoio/hugo/releases). Crunchy docs currently build with Hugo 0.55.6.

## How to set up

If:

- your project is initialized as a git repository
- you are the project owner (i.e. not working from a fork)
- this template is being added to your project **for the first time**:

1. Download this repository in ZIP format and extract the contents into your project repository.
    - If your project repository contains other source code, the contents of priv-all-doc-userguide-template should go into a `doc` subdirectory. **Change into this subdirectory before moving on to the next step.**
2. Add the [Crunchy Hugo theme](https://github.com/CrunchyData/crunchy-hugo-theme) as a git submodule:

```
git submodule add https://github.com/CrunchyData/crunchy-hugo-theme themes/crunchy-hugo-theme
```
This will create a new `themes` subdirectory. Check to see that this subdirectory contains `crunchy-hugo-theme`, and that `crunchy-hugo-theme` is not empty.

### If you are working on a fork:
 
Assuming that the upstream repo has the template and the submodule already added as in the steps above, in the `doc` subdirectory of your local dev environment you will need to run:

```
git submodule update --init
```

This will fetch the Crunchy theme data so that you can also test content edits on your local machine. (So, `/themes/crunchy-hugo-theme` should not be empty.)

## How to add content

Add and edit content in the `content` [subdirectory](./content/). Images and other assets go under the `static` subdirectory.

To learn more, check out this page on the Hugo [directory structure](https://gohugo.io/getting-started/directory-structure/).

## How to test locally

In the user guide (`doc`) root directory, run the Hugo server:

```
hugo server
```
Then, go to localhost:1313 in your browser.

## Before You Write

Ask yourself:

- Who is going to read this guide? (Is there any chance that they will be a complete beginner?)
- What does the reader need to know for everyday usage? How do you make it easy for them to find specific information?
