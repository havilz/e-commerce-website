/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{vue,js}'],
  theme: {
    extend: {
      colors: {
        primary: {
          50:  '#f0fdf4',
          100: '#dcfce7',
          400: '#4ade80',
          500: '#22c55e',
          600: '#16a34a',
          700: '#15803d',
        },
        secondary: {
          400: '#fb923c',
          500: '#f97316',
          600: '#ea580c',
        },
        accent: {
          400: '#c084fc',
          500: '#a855f7',
          600: '#9333ea',
        },
        neutral: {
          50:  '#fafafa',
          100: '#f5f5f5',
          200: '#e5e5e5',
          400: '#a3a3a3',
          600: '#525252',
          800: '#262626',
          900: '#171717',
        },
        warning: '#f59e0b',
        success: '#22c55e',
        error:   '#ef4444',
        info:    '#3b82f6',
      },
      fontFamily: {
        sans:    ['Inter', 'system-ui', 'sans-serif'],
        heading: ['"Plus Jakarta Sans"', 'system-ui', 'sans-serif'],
      },
      animation: {
        'fade-in':   'fadeIn 0.2s ease-in-out',
        'slide-up':  'slideUp 0.3s ease-out',
        'bounce-in': 'bounceIn 0.4s cubic-bezier(0.68,-0.55,0.265,1.55)',
      },
      keyframes: {
        fadeIn:   { from: { opacity: '0' },                                        to: { opacity: '1' } },
        slideUp:  { from: { transform: 'translateY(16px)', opacity: '0' },         to: { transform: 'translateY(0)', opacity: '1' } },
        bounceIn: { from: { transform: 'scale(0.8)',       opacity: '0' },         to: { transform: 'scale(1)',      opacity: '1' } },
      },
    },
  },
  plugins: [],
}
