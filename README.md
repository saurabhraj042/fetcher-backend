# { Fetcher }

This is the backend part of the web app [fetcher](https://fetcher-frontend.herokuapp.com/) made in **Golang.**

It basically scrapes this [portal](https://www.amity.edu/placement/students.asp?id=6) and gets all the notices encoded in **Json** format and exposes the api over [this](https://evening-springs-70151.herokuapp.com/) url.

For building the scraper [Colly](https://github.com/gocolly/colly) framework is used :) and deployed using **Heroku.**

The frontend of fetcher is made using ReactJs and the here is its [repository](https://github.com/saurabhraj042/fetcher-frontend).

![1633277840919.png](image/README/1633277840919.png)
