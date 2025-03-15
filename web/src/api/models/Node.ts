/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Vehicle } from './Vehicle';
import type { VehicleType } from './VehicleType';
export type Node = {
    /**
     * Уникальный идентификатор узла
     */
    id: string;
    types: Array<VehicleType>;
    vehicles?: Array<Vehicle>;
};

