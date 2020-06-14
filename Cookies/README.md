## **To delete a cookie**

#### 1 variant

`c.MaxAge = x`

Where x is a number of seconds when cookie stays alive, then it'd be deleted.

#### 2 variant

`c.MaxAge = 0`

Here '0' but it can be negative value too.
_(instantly deleting the cookie called 'c')_

### THEN

You always have to:

`http.SetCookie(w, c)`

\- and maybe you wanna do this stuff:

`http.Redirect(w, req, "/", http.StatusSeeOther`

