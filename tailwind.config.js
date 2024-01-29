/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        './pkg/templates/**/*.templ'
    ],
    plugins: [require("daisyui")],
}