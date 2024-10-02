/** @type {import('tailwindcss').Config} */

const defaultTheme = require('tailwindcss/defaultTheme');

export default {
    content: [
        "./index.html",
        "./src/**/*.{vue,js,ts,jsx,tsx}",
    ],
    theme: {
        extend: {
            boxShadow: {
                blue: "0 0 3px 3px #4285f4",
            },
            colors: {
                "brand-bg-1": "#ffffff",
                "brand-gray-1": "#dadce0",
                "brand-blue-1": "#1967d2",
                "brand-green-1": "#137333",
            },
            fontFamily: {
                sans: ['Open Sans', ...defaultTheme.fontFamily.sans],
            },
        },
    },
    plugins: [],
};