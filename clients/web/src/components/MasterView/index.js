import React from "react";

import ScrollView from "../ScrollView";

const MasterView = ({ children }) => {
    return (
        <ScrollView w={[1, 1, 1 / 2]} maxWidth={[null, null, 400]}>
            {children}
        </ScrollView>
    );
};

export default MasterView;
