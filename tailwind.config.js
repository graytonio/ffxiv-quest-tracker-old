/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        './pkg/templates/**/*.templ',
        './assets/dist/**/*'
    ],
    daisyui: {
        themes: ["dark"]
    },
    plugins: [require("daisyui")],
}