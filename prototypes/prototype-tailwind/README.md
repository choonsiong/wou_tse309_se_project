# Setup

```
[mbp2022 3:30:11:167] mkdir src dist
[mbp2022 3:30:13:167] ll
total 0
drwxr-xr-x  4 choonsiong  staff  128 Jul 18 03:30 ./
drwxr-xr-x  7 choonsiong  staff  224 Jul 18 03:29 ../
drwxr-xr-x  2 choonsiong  staff   64 Jul 18 03:30 dist/
drwxr-xr-x  2 choonsiong  staff   64 Jul 18 03:30 src/
[mbp2022 3:30:14:167] pwd
/Users/choonsiong/WOU/Semester 13,14,15 2024/github/wou_tse309_se_project/prototype-tailwind
[mbp2022 3:30:15:167] npm install -D tailwindcss

added 113 packages in 7s

29 packages are looking for funding
  run `npm fund` for details
[mbp2022 3:30:32:167] 
[mbp2022 3:30:34:167] npx tailwindcss init

Created Tailwind CSS config file: tailwind.config.js
[mbp2022 3:30:41:167] 
[mbp2022 3:30:42:167] 
[mbp2022 3:30:42:167] ls -l
total 112
drwxr-xr-x   2 choonsiong  staff     64 Jul 18 03:30 dist/
drwxr-xr-x  99 choonsiong  staff   3168 Jul 18 03:30 node_modules/
-rw-r--r--   1 choonsiong  staff  48905 Jul 18 03:30 package-lock.json
-rw-r--r--   1 choonsiong  staff     59 Jul 18 03:30 package.json
drwxr-xr-x   2 choonsiong  staff     64 Jul 18 03:30 src/
-rw-r--r--   1 choonsiong  staff    128 Jul 18 03:30 tailwind.config.js
[mbp2022 3:30:43:167] 
```

# Running

```
[mbp2022 3:36:34:167] cat tailwind.config.js 
/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./src/**/*.{html,css,js}"],
    theme: {
        extend: {},
    },
    plugins: [],
}

[mbp2022 3:36:36:167] 
[mbp2022 3:36:37:167] npx tailwindcss -i ./src/input.css -o ./dist/output.css --watch

Rebuilding...

warn - No utility classes were detected in your source files. If this is unexpected, double-check the `content` option in your Tailwind CSS configuration.
warn - https://tailwindcss.com/docs/content-configuration

Done in 86ms.
```