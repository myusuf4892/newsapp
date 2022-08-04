# News App

<!-- ABOUT THE PROJECT -->

## 💻 About The Project

News app can posting articles for media.

Features in News App :

### USERS

At users, there are features for login either user or admin, we also make CRUD for the user here


| Feature User | Endpoint | Param | JWT Token | Function |
| :---: | --- | :---: | :---: | --- |
| POST | /register  | - | NO | new account registration |
| POST | /login | - | NO | login for user |

<br>

### POSTS

At posts, we can be posting articles by user. so, you must be register and login, before posting article.

| Feature Post | Endpoint | Param | JWT Token | Function |
| :---: | --- | :---: | :---: | --- |
| POST | /posts | - | YES | create posting |
| GET | /posts | - | YES | get posts |
| PUT | /posts/id | id: post_id | YES | update post |
| DELETE | /posts/id | id: post_id | YES | delete post |
| POST | /posts/categories | - | YES | create post type |
| GET | /posts/categories | - | YES | get post type |
| PUT | /posts/categories/id | id: post_type_id | YES | update post type |
| DELETE | /posts/categories/id | id: post_type_id | YES | delete post type |

<br>

### 🛠 &nbsp;Build App & Database

![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)

![Visual Studio Code](https://img.shields.io/badge/Visual%20Studio%20Code-0078d7.svg?style=for-the-badge&logo=visual-studio-code&logoColor=white)

![MySQL](https://img.shields.io/badge/mysql-%2300f.svg?style=for-the-badge&logo=mysql&logoColor=white)

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

![JWT](https://img.shields.io/badge/JWT-black?style=for-the-badge&logo=JSON%20web%20tokens)

![Swagger](https://img.shields.io/badge/-Swagger-%23Clojure?style=for-the-badge&logo=swagger&logoColor=white)

![Postman](https://img.shields.io/badge/Postman-FF6C37?style=for-the-badge&logo=postman&logoColor=white)

## 🗃️ ERD

![ERD](./documents/news.jpg)

## Run Locally

Clone the project

```bash

  git clone https://github.com/myusuf4892/newsapp.git

```

Go to the project directory

```bash

cd newsapp

```

Run aplication

```bash

go run main.go

```

Go test aplication

```bash

chmod 777 ./test.sh
go test ./test.sh

```

## Open Api

if you want to consume our api,

here's the way !

```bash
https://app.swaggerhub.com/apis-docs/myusuf4892/newsapp/1.0#/
```

## Authors

- Muhamad Yusup

  Reach me:

[![LinkedIn](https://img.shields.io/badge/Muhamad%20Yusup-%230077B5.svg?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/muhamad-yusup-a69225234/)

[![GitHub](https://img.shields.io/badge/myusuf4892-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)](https://github.com/syawaladiyaksa15)

<p align="right">(<ahref="#top">back to top</a>)</p>

<h3>

<p align="center">:copyright: 2022 </p>

</h3>

<!-- end -->
