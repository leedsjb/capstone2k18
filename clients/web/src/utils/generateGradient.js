// sources:
// https://raw.githubusercontent.com/tobiaslins/avatar/master/src/helper.js
// https://github.com/tobiaslins/avatar/blob/master/src/image.js

import crypto from "crypto";
import Color from "color";

function djb2(str) {
    let hash = 5381;
    for (let i = 0; i < str.length; i++) {
        hash = (hash << 5) + hash + str.charCodeAt(i);
    }
    return hash;
}

function shouldChangeColor(color) {
    const rgb = color.rgb().array();
    const val = 765 - (rgb[0] + rgb[1] + rgb[2]);
    if (val < 250 || val > 700) {
        return true;
    }
    return false;
}

function hashStringToColor(str) {
    const hash = djb2(str);
    const r = (hash & 0xff0000) >> 16;
    const g = (hash & 0x00ff00) >> 8;
    const b = hash & 0x0000ff;
    return (
        "#" +
        ("0" + r.toString(16)).substr(-2) +
        ("0" + g.toString(16)).substr(-2) +
        ("0" + b.toString(16)).substr(-2)
    );
}

function getMatchingColor(firstColor) {
    let color = firstColor;
    if (color.isDark()) {
        color = color.saturate(0.3).rotate(90);
    } else {
        color = color.desaturate(0.3).rotate(90);
    }
    if (shouldChangeColor(color)) {
        color = color.rotate(-200).saturate(0.5);
    }
    return color;
}

export default function generateGradient(str) {
    const hash = crypto
        .createHash("md5")
        .update(str)
        .digest("hex");

    let firstColor = hashStringToColor(hash);
    firstColor = new Color(firstColor).saturate(0.5);

    const lightning = firstColor.hsl().color[2];
    if (lightning < 25) {
        firstColor = firstColor.lighten(3);
    }
    if (lightning > 25 && lightning < 40) {
        firstColor = firstColor.lighten(0.8);
    }
    if (lightning > 75) {
        firstColor = firstColor.darken(0.4);
    }

    return [firstColor.hex(), getMatchingColor(firstColor).hex()];
}
