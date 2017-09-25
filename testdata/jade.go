package testdata

// FRONTEND is returned from the webpage home
const FRONTEND string = `<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="utf-8">
        <link href='./main.css' rel='stylesheet'>
        <title>Title</title>
    </head>
    <body>
        <h1>Repo Finder</h1><br>
        <p>Enter github username and reponame in the input fields or</p>
        <p><b>check out some of these repos:</b></p><br>
        <table>
            <thead>
                <tr>
                    <td><b>Username</b></td>
                    <td><b>Repository name</b></td>
                    <td><b>Link</b></td>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <td>git</td>
                    <td>git</td>
                    <td><a href="/projectinfo/v1/github.com/git/git">git/git</a></td>
                </tr>
                <tr>
                    <td>npm</td>
                    <td>npm</td>
                    <td><a href="/projectinfo/v1/github.com/npm/npm">npm/npm</a></td>
                </tr>
                <tr>
                    <td>atom</td>
                    <td>atom</td>
                    <td><a href="/projectinfo/v1/github.com/atom/atom">atom/atom</a></td>
                </tr>
            </tbody>
        </table>
        <form class="form" method="get" style="display: flex; flex-direction: column;">
            <input class="input" name="userName" type="text" placeholder="username" style="border: 1px solid grey; background: #fff; border-radius: 4px; color: #333; box-shadow: rgba(255,255,255,0.4) 0 1px 0, inset rgba(000,000,000,0.7) 0 0px 0px; padding: 8px; margin-bottom: 10px;">
            <input class="input" name="repoName" type="text" placeholder="repository name" style="border: 1px solid grey; background: #fff; border-radius: 4px; color: #333; box-shadow: rgba(255,255,255,0.4) 0 1px 0, inset rgba(000,000,000,0.7) 0 0px 0px; padding: 8px; margin-bottom: 10px;">
            <button class="button" type="submit" style="background: #3498db; background-image: linear-gradient(to bottom, #3498db, #3498db); border-radius: 6px; color: #ffffff; font-size: 20px; padding: 10px 20px 10px 20px; border:none">Submit</button>
        </form>
    </body>
</html>`
