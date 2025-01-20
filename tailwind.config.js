/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./layouts/*.gohtml",
    "./template/*.gohtml",
    "./node_modules/flowbite/**/*.js",
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('flowbite/plugin')
  ]
}
