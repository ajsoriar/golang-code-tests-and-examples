
    /*
        Run this file this way:

        go run example1.go
    */

    package main

    import (
        _ "github.com/go-sql-driver/mysql"
        "database/sql"
        "fmt"
    )

    /*
        // This is data about my local database

        $user = 'root';
        $password = 'root';
        $db = 'inventory'; // andres-db
        $host = 'localhost';
        $port = 8889;

        $link = mysql_connect(
           "$host:$port", 
           $user, 
           $password
        );
        $db_selected = mysql_select_db(
           $db, 
           $link
        );
    */

    func main() {

        // db, err := sql.Open("mysql", "astaxie:astaxie@/test?charset=utf8")
        // checkErr(err)

        /*
            Check: http://go-database-sql.org/accessing.html
        */

        // func main() {
        //     db, err := sql.Open("mysql",
        //         "user:password@tcp(127.0.0.1:3306)/hello")
        //     if err != nil {
        //         log.Fatal(err)
        //     }
        //     defer db.Close()
        // }

        db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/andres-db")

        // if err != nil {
        //     log.Fatal(err)
        // }

        err = db.Ping()
        if err != nil {
            // do something here
            fmt.Printf("Some kind of error: %v\n", err)
        }

        //defer db.Close()

        // insert

        /*
        stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
        checkErr(err)

        res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
        checkErr(err)

        id, err := res.LastInsertId()
        checkErr(err)

        fmt.Println(id)
        */

        // update

        /*
        stmt, err = db.Prepare("update userinfo set username=? where uid=?")
        checkErr(err)

        res, err = stmt.Exec("astaxieupdate", id)
        checkErr(err)

        affect, err := res.RowsAffected()
        checkErr(err)

        fmt.Println(affect)
        */

        // query

        /*
        rows, err := db.Query("SELECT * FROM userinfo")
        checkErr(err)

        for rows.Next() {
            var uid int
            var username string
            var department string
            var created string
            err = rows.Scan(&uid, &username, &department, &created)
            checkErr(err)
            fmt.Println(uid)
            fmt.Println(username)
            fmt.Println(department)
            fmt.Println(created)
        }
        */

        rows, err := db.Query("SELECT * FROM users")
        checkErr(err)

        fmt.Printf("Reading the table here:\n")

        for rows.Next() {
            var id int
            var username string
            err = rows.Scan(&id, &username)
            checkErr(err)
            fmt.Println(id)
            fmt.Println(username)
        }

        // delete

        /*
        stmt, err = db.Prepare("delete from userinfo where uid=?")
        checkErr(err)

        res, err = stmt.Exec(id)
        checkErr(err)

        affect, err = res.RowsAffected()
        checkErr(err)

        fmt.Println(affect)
        */

        db.Close()
    }

    func checkErr(err error) {
        if err != nil {
            //panic(err)

            fmt.Printf("Some kind of error: %v\n", err)
        }
    }