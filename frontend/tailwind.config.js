import flowbite from "flowbite-react/tailwind";

/** @type {import('tailwindcss').Config} */
export default {
  content: ["./src/**/*.{html,js,jsx}", './src/components/**/*.jsx', './src/pages/**/*.jsx', flowbite.content()],
  theme: {
    extend: {
      spacing: {
        '128': '32rem',
      },
      color: {
        'graytah': '#eaebed'
      }
    },
  },
  plugins: [flowbite.plugin()],
}

