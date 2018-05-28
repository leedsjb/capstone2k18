const colors = {
    airlift1: "#E4002B",
    airlift2: "#ED002B",
    black1: "#141414",
    black2: "#292929",
    black3: "#3D3D3D",
    gray1: "#A3A3A3",
    gray2: "#8F8F8F",
    gray3: "#EBEBEB",
    gray4: "#FAFAFA"
};

const boxShadows = {
    low: `0 2px 8px ${colors.gray3}`,
    mid: `0 4px 12px ${colors.gray3}`,
    high: `0 8px 16px ${colors.gray3}`
};

const borders = {
    light: `1px solid ${colors.gray3}`
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
    boxShadows,
    borders,
    space,
    fontSizes,
    fontWeights,
    fonts,
    colors,
    radii,
    shadows
};

export default theme;
