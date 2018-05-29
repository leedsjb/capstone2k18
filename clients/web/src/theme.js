const colors = {
    airlift1: "#E4002B",
    airlift2: "#ED002B",
    black1: "#141414",
    black2: "#292929",
    black3: "#3D3D3D",
    gray1: "#A3A3A3",
    gray2: "#8F8F8F",
    gray3: "#DBDBDB",
    gray4: "#E0E0E0",
    gray5: "#EBEBEB",
    gray6: "#F5F5F5",
    gray7: "#FAFAFA",
    green1: "#008b00",
    green2: "#f2fde4",
    orange1: "#f47100",
    orange2: "#fff2df",
    purple: "#7e3ff2",
    blue: "#4a26fd",
    pink1: "#ef0078",
    pink2: "#fbe2f0",
    white: "#FFFFFF"
};

const boxShadows = {
    low: `0 2px 8px rgba(0, 0, 0, 0.08)`,
    lowRight: `2px 2px 8px rgba(0, 0, 0, 0.08)`,
    mid: `0 4px 12px rgba(0, 0, 0, 0.08)`,
    high: `0 8px 16px rgba(0, 0, 0, 0.08)`
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
    space,
    fontSizes,
    fontWeights,
    fonts,
    colors,
    radii,
    shadows
};

export default theme;
