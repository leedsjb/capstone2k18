import React from "react";
import { Helmet } from "react-helmet";

import FlexFullHeight from "../../components/FlexFullHeight";
import TitleBar from "../../components/TitleBar";
import TabBar from "../../components/TabBar";
import ScrollView from "../../components/ScrollView";
import Accordion from "../../components/Accordion";
import AccordionSection from "../../components/AccordionSection";
import Text from "../../components/Text";

const MissionDetailPage = () => {
    return (
        <FlexFullHeight flexDirection="column">
            <Helmet>
                <title>Missions</title>
            </Helmet>
            <TitleBar back backPath="/missions" />
            <ScrollView>
                <Accordion>
                    <AccordionSection title="Radio report">
                        <Text>123</Text>
                    </AccordionSection>
                    <AccordionSection title="Assigned crew">
                        <Text>123</Text>
                    </AccordionSection>
                    <AccordionSection title="Requestor">
                        <Text>123</Text>
                    </AccordionSection>
                    <AccordionSection title="Receiver">
                        <Text>123</Text>
                    </AccordionSection>
                </Accordion>
            </ScrollView>
            <TabBar />
        </FlexFullHeight>
    );
};

export default MissionDetailPage;
