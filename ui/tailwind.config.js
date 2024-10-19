/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{html,js,tsx}"],
  theme: {
    extend: {
      colors: {
        main: {
          normal: '#10B981', // emerald-500
          hover: '#059669', // emerald-600
        }, 
        accent: {
          normal: '#64748b', // slate-500
          hover: '#475569', // slate-600
        },
        background: {
          main: '#000000', // black
          accent: {
            normal: '#18181b', // zinc-900
            hover: '#3f3f46', // zinc-700
          }
        },
      },
    },
  },
  plugins: [],
  darkMode: 'selector',
}