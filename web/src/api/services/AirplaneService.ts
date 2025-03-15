/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class AirplaneService {
    /**
     * Получение свободного места парковки для самолета
     * В зависимости от загрузки парковок отдает нужный узел
     * @param id ID самолета
     * @returns any Id узла парковочного места
     * @throws ApiError
     */
    public static airplaneGetParkingSpot(
        id: string,
    ): CancelablePromise<{
        /**
         * ID узла
         */
        nodeId: string;
    }> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/airplane/{id}/parking',
            path: {
                'id': id,
            },
            errors: {
                400: `Неверные данные запроса`,
                409: `Нет свободного парковочного места для самолета`,
            },
        });
    }
    /**
     * Фиксация вылета самолета
     * Удаляется самолет с карты
     * @param id ID самолета
     * @returns any Самолет улетел
     * @throws ApiError
     */
    public static airplaneTakeOff(
        id: string,
    ): CancelablePromise<any> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/airplane/{id}/take-off',
            path: {
                'id': id,
            },
            errors: {
                404: `Самолет не найден на ВПП`,
            },
        });
    }
}
