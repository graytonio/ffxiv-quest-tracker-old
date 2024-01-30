/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        './pkg/templates/**/*.templ'
    ],
    daisyui: {
        themes: ["dark"]
    },
    plugins: [require("daisyui")],
}