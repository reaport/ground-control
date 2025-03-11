/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { AirportMap } from '../models/AirportMap';
import type { AirportMapConfig } from '../models/AirportMapConfig';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class MapService {
    /**
     * Получить карту аэропорта
     * Возвращает полную карту аэропорта в виде графа.
     * @returns AirportMap Успешный запрос
     * @throws ApiError
     */
    public static mapGetAirportMap(): CancelablePromise<AirportMap> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/map',
        });
    }
    /**
     * Обновить карту аэропорта
     * Обновляет карту аэропорта.
     * @param requestBody
     * @returns any Успешный запрос
     * @throws ApiError
     */
    public static mapUpdateAirportMap(
        requestBody: AirportMap,
    ): CancelablePromise<any> {
        return __request(OpenAPI, {
            method: 'PUT',
            url: '/map',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Неверные данные`,
            },
        });
    }
    /**
     * Возвращает карту к исходному состоянию
     * Возвращает карту к исходному состоянию
     * @returns any Успешный запрос
     * @throws ApiError
     */
    public static mapRefreshAirportMap(): CancelablePromise<any> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/map/refresh',
        });
    }
    /**
     * Получить конфигурацию карты аэропорта
     * Возвращает конфигурацию аэропорта.
     * @returns AirportMapConfig Успешный запрос
     * @throws ApiError
     */
    public static mapGetAirportMapConfig(): CancelablePromise<AirportMapConfig> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/map/config',
        });
    }
}
