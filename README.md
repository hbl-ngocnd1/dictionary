[![tests](https://github.com/hbl-ngocnd1/dictionary/actions/workflows/tests.yml/badge.svg)](https://github.com/hbl-ngocnd1/dictionary/actions/workflows/tests.yml)
[![Deploy to Amazon ECS](https://github.com/hbl-ngocnd1/dictionary/actions/workflows/aws.yml/badge.svg)](https://github.com/hbl-ngocnd1/dictionary/actions/workflows/aws.yml)
[![Coverage Status](https://coveralls.io/repos/github/hbl-ngocnd1/dictionary/badge.svg?branch=main)](https://coveralls.io/github/hbl-ngocnd1/dictionary?branch=main)

## Overview

1. Dictionary execution get list word from [japanesetest4you](https://japanesetest4you.com/jlpt-n1-vocabulary-list/) with level is n1~n5

2. With one word was show in the main page, use one goroutine to concurrency get detail info in look like [detail page](https://japanesetest4you.com/flashcard/%e8%b5%a4%e5%ad%97-akaji/)

3. After get detail info, use [google translate api](https://pkg.go.dev/cloud.google.com/go/translate) to bulk and concurrency translated english mean to vietnamese mean

4. Save data to global variable(cache in memory) after first request from server start, data was get from cache

## Code Architecture 

Use Clean Architecture reference [this blogs](https://qiita.com/ogady/items/34aae1b2af3080e0fec4)

## Database

In dictionary feature not use database.

But in the future can be use elastic search or self build in search engine like [this](https://zenn.dev/yukiyada/articles/7e2c67d8406f0d) to search future 

## Prerequisites

You'll need the following:
* [Git](https://git-scm.com/downloads)
* [Go](https://golang.org/dl/)

## 1. Clone the app

Now you're ready to start working with the simple Go *hello world* app. Clone the repository and change to the directory where the sample app is located.
  ```
git clone https://github.com/hbl-ngocnd1/dictionary
cd dictionary
  ```

Peruse the files in the *dictionary* directory to familiarize yourself with the contents.

## 2. Run the app locally use [air](https://github.com/cosmtrek/air)
Create .env file
```cmd
CLOUDANT_URL=dasdsa
GOOGLE_APPLICATION_API_KEY=sss
DEBUG=true
SYNC_PASS=123456
```
Build and run the app.
  ```
air
  ```

View your app at: http://localhost:8080