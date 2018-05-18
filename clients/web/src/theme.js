const colors = {
    wireframe: "#d8d8d8",
    primary: "#3E50C2",
    primaryLight: "#4052C7",
    secondary: "#E31B3D",
    secondaryLight: "#DB4861",
    border: "#D4D5D6",
    gray: "#F7F8FC"
};

const breakpoints = [32, 48, 64, 80].map(n => n + "em");

let space = [];

for (let i = 0; i < 37; i++) {
    space.push(i * 4);
}

const fontSizes = [12, 14, 16, 20, 24, 32, 48, 64, 72, 96];

const fontWeights = {
    normal: 400,
    bold: 700
};

const radii = [0, 2, 4];

const fonts = {
    0: "system-ui, sans-serif",
    sans: "system-ui, sans-serif",
    mono: '"SF Mono", "Roboto Mono", Menlo, monospace'
};

const shadows = [
    "none",
    `inset 0 0 0 1px ${colors.wireframe}`,
    `inset 0 0 0 1px ${colors.wireframe}, 0 0 4px ${colors.wireframe}`
];

const theme = {
    breakpoints,
    space,
    fontSizes,
    fontWeights,
    fonts,
    colors,
    radii,
    shadows
};

export default theme;
