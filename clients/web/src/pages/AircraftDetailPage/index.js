import React from "react";
import { Helmet } from "react-helmet";

import FlexFillVH from "../../components/FlexFillVH";
import TitleBar from "../../components/TitleBar";
import TabBar from "../../components/TabBar";
import ScrollView from "../../components/ScrollView";
import Accordion from "../../components/Accordion";
import AccordionSection from "../../components/AccordionSection";
import Text from "../../components/Text";

const AircraftDetailPage = () => {
    let backPath =
        new URLSearchParams(window.location.search).get("source") === "map"
            ? "/aircraft/map"
            : "/aircraft";
    console.log(backPath);
    return (
        <FlexFillVH flexDirection="column">
            <Helmet>
                <title>Missions</title>
            </Helmet>
            <TitleBar back backPath={backPath} />
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
        </FlexFillVH>
    );
};

export default AircraftDetailPage;
