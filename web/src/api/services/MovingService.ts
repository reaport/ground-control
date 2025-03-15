/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { VehicleType } from '../models/VehicleType';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class MovingService {
    /**
     * Запросить маршрут
     * Запрашивает маршрут из точки А в точку Б.
     * @param requestBody
     * @returns string Успешный запрос
     * @throws ApiError
     */
    public static movingGetRoute(
        requestBody: {
            /**
             * ID начального узла
             */
            from: string;
            /**
             * ID конечного узла
             */
            to: string;
            /**
             * Тип транспорта
             */
            type: VehicleType;
        },
    ): CancelablePromise<Array<string>> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/route',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                404: `Маршрут не найден`,
            },
        });
    }
    /**
     * Запросить разрешение на перемещение
     * Запрашивает разрешение на перемещение из одного узла в другой.
     * @param requestBody
     * @returns any Разрешение получено
     * @throws ApiError
     */
    public static movingRequestMove(
        requestBody: {
            /**
             * ID транспорта
             */
            vehicleId: string;
            vehicleType: VehicleType;
            /**
             * ID текущего узла
             */
            from: string;
            /**
             * ID следующего узла
             */
            to: string;
            /**
             * ID самолета, который следует за follow-me
             */
            withAirplane?: string;
        },
    ): CancelablePromise<{
        /**
         * Расстояние до следующего узла
         */
        distance: number;
    }> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/move',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Неверные данные запроса`,
                403: `Перемещение запрещено`,
                404: `Один из узлов не найден`,
                409: `Узел занят`,
            },
        });
    }
    /**
     * Уведомить о прибытии в узел
     * Уведомляет вышку о прибытии транспорта в узел.
     * @param requestBody
     * @returns any Уведомление успешно обработано
     * @throws ApiError
     */
    public static movingNotifyArrival(
        requestBody: {
            /**
             * ID транспорта
             */
            vehicleId: string;
            vehicleType: VehicleType;
            /**
             * ID узла
             */
            nodeId: string;
        },
    ): CancelablePromise<any> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/arrived',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Неверные данные запроса`,
                404: `Узел не найден`,
            },
        });
    }
    /**
     * Регистрация транспорта на карте
     * В зависимости от типа транспорта отдает нужную начальную точку и id
     * @param type Тип транспорта
     * @returns any Id узла начальной точки
     * @throws ApiError
     */
    public static movingRegisterVehicle(
        type: VehicleType,
    ): CancelablePromise<{
        /**
         * ID узла
         */
        garrageNodeId: string;
        /**
         * ID транспорта
         */
        vehicleId: string;
        /**
         * Id узлов парковочных мест для обслуживания самолетов (парковка\_самолета:парковка\_сервисной\_машинки)
         */
        serviceSpots: Record<string, string>;
    }> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/register-vehicle/{type}',
            path: {
                'type': type,
            },
            errors: {
                400: `Неверные данные запроса`,
                409: `Взлетно-посадочная полоса занята`,
            },
        });
    }
}
