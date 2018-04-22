import React from "react";
import { Helmet } from "react-helmet";

import TitleBar from "../../components/TitleBar";
import Container from "../../components/Container";
import TabBar from "../../components/TabBar";
import FlexFillVH from "../../components/FlexFillVH";
import ScrollView from "../../components/ScrollView";

import ResourcesProvider from "../../containers/ResourcesProvider";

const ResourcesPage = () => {
    return (
        <FlexFillVH flexDirection="column">
            <Helmet>
                <title>Resources</title>
            </Helmet>

            <TitleBar title="Resources" />

            <ScrollView>
                <Container>
                    <div>Resources</div>
                    <ResourcesProvider
                        render={resources => {
                            return <div />;
                        }}
                    />
                </Container>
            </ScrollView>

            <TabBar />
        </FlexFillVH>
    );
};

export default ResourcesPage;
