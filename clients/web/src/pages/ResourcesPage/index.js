import React from "react";
import { Helmet } from "react-helmet";

import TitleBar from "../../components/TitleBar";
import Container from "../../components/Container";
import TabBar from "../../components/TabBar";
import FlexFullHeight from "../../components/FlexFullHeight";
import ScrollView from "../../components/ScrollView";

import ResourcesProvider from "../../containers/ResourcesProvider";

const ResourcesPage = () => {
    return (
        <FlexFullHeight flexDirection="column">
            <Helmet>
                <title>Resources</title>
            </Helmet>

            <TitleBar title="Resources" />

            <ScrollView>
                <Container>
                    <div>Resources</div>
                    <ResourcesProvider
                        render={resources => {
                            console.log(resources);
                            return <div />;
                        }}
                    />
                </Container>
            </ScrollView>

            <TabBar />
        </FlexFullHeight>
    );
};

export default ResourcesPage;
