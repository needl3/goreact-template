const defaultTheme = require("tailwindcss/defaultTheme");

module.exports = {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
    "./public/views/*.html",
  ],
  theme: {
    screens: {
      sm: "768px",

      smmd: "792px",

      md: "1024px",

      lg: "1280px",

      xl: "1536px",
    },
    extend: {
      fontFamily: {
        sans: ["Inter var", ...defaultTheme.fontFamily.sans],
      },
      backgroundImage: {
        "profile-cover": "url('/cover.png')",
      },
      colors: {
        disabled: "#5F75CD",
        primary: {
          900: "#0A175B", // Original color
          800: "#1D2A7F",
          700: "#3043A2",
          600: "#475FBF",
          500: "#5C7BDB",
          400: "#8297E7",
          300: "#A2B2F2",
          200: "#C0CDEE",
          100: "#DEE9FA",
          light: "#0A5FEE",
          dark: "#0A175B",
          default: "#0A5FEE",
        },
        secondary: {
          default: "#34469C",
        },
        success: {
          900: "#349400",
          800: "#58b300",
          700: "#7bc90c",
          600: "#9de019",
          500: "#baff26", // Original color
          400: "#d6ff63",
          300: "#f1ff9f",
          200: "#fdffd5",
          100: "#fffffa",
          light: "#fffffa",
          dark: "#349400",
          default: "#baff26",
        },
        info: {
          900: "#9dcfd0",
          800: "#afd8d8",
          700: "#c2e1e0",
          600: "#d4e9e8",
          500: "#e7f2f1", // Original color
          400: "#c3e4e1",
          300: "#9fd6d0",
          200: "#7bc7bf",
          100: "#57b9ae",
          light: "#57b9ae",
          dark: "#9dcfd0",
          default: "#e7f2f1",
        },
        warning: {
          900: "#c87c00",
          800: "#d38d00",
          700: "#dd9d00",
          600: "#e8ae00",
          500: "#f2be02", // Original color
          400: "#f8d654",
          300: "#fadd90",
          200: "#fde3c6",
          100: "#fef1e8",
          light: "#fef1e8",
          dark: "#c87c00",
          default: "#f2be02", // Original color
        },
        danger: {
          900: "#b11b00",
          800: "#bb2f0c",
          700: "#c54319",
          600: "#cf5724",
          500: "#e25724", // Original color
          400: "#ea6c3a",
          300: "#f28050",
          200: "#f89566",
          100: "#ffb082",
          light: "#ffb082",
          dark: "#b11b00",
          default: "#e25724", // Original color
        },
      },
    },
  },
  plugins: [
    require("@tailwindcss/aspect-ratio"),
    require("@tailwindcss/typography"),
    require("@tailwindcss/container-queries"),
  ],
};
