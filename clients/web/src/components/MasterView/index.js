import React from "react";
import { Flex } from "grid-styled";

import NavBarItem from "../NavBarItem";
import ScrollView from "../ScrollView";

const MasterView = ({ children }) => {
    return (
        <ScrollView w={[1, 1, 1 / 2]} maxWidth={[null, null, 400]}>
            <Flex>
                <NavBarItem title="Aircraft" path="/aircraft" />
                <NavBarItem title="People" path="/people" />
            </Flex>

            {children}
        </ScrollView>
    );
};

export default MasterView;
