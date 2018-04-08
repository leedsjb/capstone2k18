import React, { Component } from "react";
import { connect } from "react-redux";

import { fetchAircraft } from "../../actions/missions/actions";

class AircraftProvider extends Component {
    componentDidMount() {
        this.props.fetchAircraft();
    }

    render() {
        return this.props.render({ aircraft: this.props.aircraft });
    }
}

const mapStateToProps = ({ aircraft }) => {
    return {
        aircraft
    };
};

const mapDispatchToProps = {
    fetchAircraft
};

export default connect(mapStateToProps, mapDispatchToProps)(AircraftProvider);
