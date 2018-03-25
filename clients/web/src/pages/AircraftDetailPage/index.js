import React from "react";
import { Helmet } from "react-helmet";

import FlexFullHeight from "../../components/FlexFullHeight";
import TitleBar from "../../components/TitleBar";
import TabBar from "../../components/TabBar";
import ScrollView from "../../components/ScrollView";
import Accordion from "../../components/Accordion";
import AccordionSection from "../../components/AccordionSection";
import Text from "../../components/Text";

const AircraftDetailPage = () => {
    return (
        <FlexFullHeight flexDirection="column">
            <Helmet>
                <title>Missions</title>
            </Helmet>
            <TitleBar back backPath="/aircraft" />
            <ScrollView>
                <Accordion>
                    <AccordionSection title="Crew">
                        <Text>123</Text>
                    </AccordionSection>
                    <AccordionSection title="Level of care">
                        <Text>123</Text>
                    </AccordionSection>
                </Accordion>
            </ScrollView>
            <TabBar />
        </FlexFullHeight>
    );
};

export default AircraftDetailPage;
