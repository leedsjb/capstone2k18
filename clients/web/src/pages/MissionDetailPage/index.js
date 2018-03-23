import React from "react";
import { Helmet } from "react-helmet";

import FlexFullHeight from "../../components/FlexFullHeight";
import TitleBar from "../../components/TitleBar";
import TabBar from "../../components/TabBar";
import ScrollView from "../../components/ScrollView";

const MissionDetailPage = () => {
    return (
        <FlexFullHeight flexDirection="column">
            <Helmet>
                <title>Missions</title>
            </Helmet>
            <TitleBar />
            <ScrollView>MissionDetailPage</ScrollView>
            <TabBar />
        </FlexFullHeight>
    );
};

export default MissionDetailPage;
