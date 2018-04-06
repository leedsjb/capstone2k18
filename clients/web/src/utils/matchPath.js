import { matchPath } from "react-router";

export default function(currentPath, path) {
    return matchPath(currentPath, {
        path,
        exact: false,
        strict: false
    });
}
