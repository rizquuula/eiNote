/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{html,js,tsx}"],
  theme: {
    extend: {
      fontSize: {
        'h1': ['2rem', { lineHeight: '2.5rem' }], // Set specific font size and line height
        'h2': ['1.75rem', { lineHeight: '2.25rem' }],
        'h3': ['1.5rem', { lineHeight: '2rem' }],
        'h4': ['1.25rem', { lineHeight: '1.75rem' }],
        'h5': ['1rem', { lineHeight: '1.5rem' }],
        'h6': ['0.875rem', { lineHeight: '1.25rem' }],
      },
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