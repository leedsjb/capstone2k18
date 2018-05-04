import React from "react";
import { Helmet } from "react-helmet";

import Accordion from "../../components/Accordion";
import AccordionSection from "../../components/AccordionSection";
import FlexFillVH from "../../components/FlexFillVH";
import InsetMapView from "../../components/InsetMapView";
import ScrollView from "../../components/ScrollView";
import TabBar from "../../components/TabBar";
import Text from "../../components/Text";
import TitleBar from "../../components/TitleBar";

const AircraftDetailPage = ({ match }) => {
    let backPath =
        new URLSearchParams(window.location.search).get("source") === "map"
            ? `/aircraft/map/${match.params.id}`
            : "/aircraft";
    return (
        <FlexFillVH flexDirection="column">
            <Helmet>
                <title>Missions</title>
            </Helmet>
            <TitleBar back backPath={backPath} />
            <ScrollView>
                <InsetMapView id={match.params.id} />
                <Accordion>
                    <AccordionSection title="Radio Report">
                        <Text>123</Text>
                    </AccordionSection>
                    <AccordionSection title="Assigned Crew">
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
        </FlexFillVH>
    );
};

export default AircraftDetailPage;
