## Introduction

As I want to explore more into the Go ecosystem, I want to use it to create my website using go and its template to create one instead of relying on something like [Hugo](https://gohugo.io).

In this poste, I will describe how my static blog generator is built.


## Motivation

I'm not good with Frontend stuffs, I don't know much about React, Svelte... and I feel like it's overkill anyway. My motivation is to keep things as simple as possible, and even Javascript is mostly not needed. 

The site should be really simple: a landing page which lists all my posts, an About section, and my resume (why not). I want to keep the feature to the minimum, so naturally the only one thing I really need to work on is how I write my blog.

## Concept

I decided to write everything in markdown and keep it inside my github repository or in another repository (I think keep it separately is better, will work on it soon). Markdown is good enough, and each folder represents the URL of an article.

We also need a metadata for each article to represent on the article listing sections. For now it would look like this:


```yml 
title: test 2
short: "Welcome to my blog! I'm excited to share my thoughts and experiences with you."
date: 09-03-2024
tags:
  - golang
  - go
  - templ
  - air
```

With theses informations, it should be enough for our listing sections to work on without the need to open articles.

By doing so, we can define our data model for each article:

```go
type ArticleMetadata struct {
	Title            string
	Link             string
	Date             string
	ShortDescription string
	Tags             []string
}

type Article struct {
	ID      string
	Content []byte
	Meta    ArticleMetadata
}
```

## Tech 

### On coding
Ofcourse go is the main language for this blog generator, with its dependences:


- [templ](https://templ.guide/): I heard a lot about this HTML builder recently and I really want to try. To be honest the ```html\template```is good enough for me but we want the new shiny thing.
- [Gomarkdown](https://github.com/gomarkdown/markdown): I'm not reinventing the wheels on how to handle markdown content. The libary is enough and also we can manipulate it to display content however we want
- [Air](https://github.com/cosmtrek/air): The needed live reload for go apps, really necessary as we build front end.  

### On deployment

Initally I want to use ``Aws lambda`` or ``Netlify``for this kind of project, it fits perfectly. 

But as the same time I have a ```Digital droplet``` laying around so why not? The thing costs 5$/month and you don't have to worry about the unexpected raising cost!

So the deployment is really straight forward.I use github to host my code so naturally github actions for deployment. Everytime we push something to the main branch, we just create a new docker registry for it:

```yml
docker-build:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v4
      - name: Set up QEMU
        uses: docker/setup-qemu-action@v3
      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v3
      - name: Login to Docker Hub
        uses: docker/login-action@v3
        with:
          username: ${{ secrets.DOCKER_USERNAME }}
          password: ${{ secrets.DOCKER_TOKEN }}
      - name: Build and push
        uses: docker/build-push-action@v5
        with:
          context: .
          platforms: linux/amd64,linux/arm64
          push: true
          tags: schierke/personal-site:latest
```

Then we ssh to our droplet, then remove old docker file, pull new one and start it.

```yml
  deploy-to-droplet:
    needs: docker-build
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v4

      - name: Install SSH client
        run: sudo apt-get install -y openssh-client

      - name: SSH into Droplet, pull image, and run container
        uses: appleboy/ssh-action@master
        with:
          host: ${{ secrets.DROPLET_SSH_HOST }}
          username: ${{ secrets.DROPLET_SSH_USERNAME }}
          key: ${{ secrets.DROPLET_SSH_PRIVATE_KEY }}
          port: ${{ secrets.DROPLET_SSH_PORT }}
          script: |
            cd /home/app/personal-site
            docker pull schierke/personal-site
            docker stop personal-site || true
            docker rm personal-site || true
            docker run -d --name personal-site --network host -p 5000:5000 --env-file /home/app/personal-site/.env schierke/personal-site
```

And that's it. We have some downtime between each deployment but honestly it's not a big deal at all; it's not like I have any reader.

## Conclusion

I hope you find my blog informative and engaging. Stay tuned for more exciting articles coming soon!

Happy reading!

