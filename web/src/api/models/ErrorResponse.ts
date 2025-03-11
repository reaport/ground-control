/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export type ErrorResponse = {
    code: ErrorResponse.code;
    /**
     * Описание ошибки
     */
    message?: string;
};
export namespace ErrorResponse {
    export enum code {
        VEHICLE_NOT_FOUND_IN_NODE = 'VEHICLE_NOT_FOUND_IN_NODE',
        EDGE_NOT_FOUND = 'EDGE_NOT_FOUND',
        MAP_HAS_VEHICLES = 'MAP_HAS_VEHICLES',
    }
}

