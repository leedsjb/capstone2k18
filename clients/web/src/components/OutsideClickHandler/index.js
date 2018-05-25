import React, { Component } from "react";

class OutsideClickHandler extends Component {
    constructor(props) {
        super(props);

        this.setWrapperRef = this.setWrapperRef.bind(this);
        this.handleClickOutside = this.handleClickOutside.bind(this);
    }

    componentDidMount() {
        document.addEventListener("mousedown", this.handleClickOutside);
    }

    componentWillUnmount() {
        document.addEventListener("mousedown", this.handleClickOutside);
    }

    setWrapperRef(node) {
        this.wrapperRef = node;
    }

    handleClickOutside(event) {
        if (this.wrapperRef && !this.wrapperRef.contains(event.target)) {
            this.props.handleClickOutside();
        }
    }

    render() {
        return (
            <div
                ref={this.setWrapperRef}
                style={{ display: "flex", flexDirection: "column", flex: 1 }}
            >
                {this.props.children}
            </div>
        );
    }
}

export default OutsideClickHandler;
