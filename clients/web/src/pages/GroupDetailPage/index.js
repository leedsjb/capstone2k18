import React from "react";
import { Helmet } from "react-helmet";

import FlexFillVH from "../../components/FlexFillVH";
import Toolbar from "../../components/Toolbar";
import SearchBox from "../../components/SearchBox";
import TitleBar from "../../components/TitleBar";
import TabBar from "../../components/TabBar";
import ScrollView from "../../components/ScrollView";

const GroupDetailPage = () => {
    return (
        <FlexFillVH flexDirection="column">
            <Helmet>
                <title />
            </Helmet>
            <TitleBar title="Aircraft" />
            <Toolbar>Test</Toolbar>
            <ScrollView>
                <SearchBox />
                Aircraft page
            </ScrollView>
            <TabBar />
        </FlexFillVH>
    );
};

export default GroupDetailPage;
