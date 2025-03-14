/* eslint-disable @typescript-eslint/no-explicit-any */
import {
  AirplaneService,
  AirportMap,
  AirportMapConfig,
  MapService,
  MovingService,
  VehicleType,
} from "@/api";
import { useMutation, UseMutationResult } from "@tanstack/react-query";
import { useState } from "react";
import { useSonner } from "sonner";

type UseFetchFunctionsResult = {
  getMapMutation: UseMutationResult<AirportMap, Error, void, unknown>;
  updateMapMutation: UseMutationResult<any, Error, AirportMap, unknown>;
  refreshMapMutation: UseMutationResult<any, Error, void, unknown>;
  registerVehicleMutation: UseMutationResult<
    {
      garrageNodeId: string;
      vehicleId: string;
      serviceSpots: Record<string, string>;
    },
    Error,
    VehicleType,
    unknown
  >;
  getRouteMutation: UseMutationResult<
    string[],
    Error,
    {
      from: string;
      to: string;
      type: VehicleType;
    },
    unknown
  >;
  requestMoveMutation: UseMutationResult<
    {
      distance: number;
    },
    Error,
    {
      vehicleId: string;
      vehicleType: VehicleType;
      from: string;
      to: string;
    },
    unknown
  >;
  notifyArrivalMutation: UseMutationResult<
    any,
    Error,
    {
      vehicleId: string;
      vehicleType: VehicleType;
      nodeId: string;
    },
    unknown
  >;
  getAirplaneParkingMutation: UseMutationResult<
    {
      nodeId: string;
    },
    Error,
    string,
    unknown
  >;
  getConfigMutation: UseMutationResult<AirportMapConfig, Error, void, unknown>;
  takeOffMutation: UseMutationResult<any, Error, string, unknown>;
  result: string | null;
};

export const useFetchFunctions = (): UseFetchFunctionsResult => {
  const [result, setResult] = useState<string | null>(" ");
  const { toasts } = useSonner();

  const getMapMutation = useMutation({
    mutationFn: () => MapService.mapGetAirportMap(),
    onSuccess: (data) => setResult(JSON.stringify(data, null, 2)),
    onError: (error) => setResult(`Ошибка: ${error}`),
  });

  const updateMapMutation = useMutation({
    mutationFn: (data: AirportMap) => MapService.mapUpdateAirportMap(data),
    onSuccess: () => setResult(null),
    onError: (error) => setResult(`Ошибка: ${error}`),
  });

  const getRouteMutation = useMutation({
    mutationFn: (data: { from: string; to: string; type: VehicleType }) =>
      MovingService.movingGetRoute(data),
    onSuccess: (data) => setResult(JSON.stringify(data, null, 2)),
    onError: (error) => setResult(`Ошибка: ${error}`),
  });

  const requestMoveMutation = useMutation({
    mutationFn: (data: {
      vehicleId: string;
      vehicleType: VehicleType;
      from: string;
      to: string;
    }) => MovingService.movingRequestMove(data),
    onSuccess: (data) => setResult(JSON.stringify(data, null, 2)),
    onError: (error) => setResult(`Ошибка: ${error.message}`),
  });

  const notifyArrivalMutation = useMutation({
    mutationFn: (data: {
      vehicleId: string;
      vehicleType: VehicleType;
      nodeId: string;
    }) => MovingService.movingNotifyArrival(data),
    onSuccess: () => setResult(null),
    onError: (error) => setResult(`Ошибка: ${error.message}`),
  });

  const registerVehicleMutation = useMutation({
    mutationFn: (type: VehicleType) =>
      MovingService.movingRegisterVehicle(type),
    onSuccess: (data) => setResult(JSON.stringify(data, null, 2)),
    onError: (error) => setResult(`Ошибка: ${error.message}`),
  });

  const takeOffMutation = useMutation({
    mutationFn: (id: string) => AirplaneService.airplaneTakeOff(id),
    onSuccess: () => setResult(null),
    onError: (error) => setResult(`Ошибка: ${error.message}`),
  });

  const refreshMapMutation = useMutation({
    mutationFn: () => MapService.mapRefreshAirportMap(),
    onSuccess: () => setResult(null),
    onError: (error) => setResult(`Ошибка: ${error.message}`),
  });

  const getAirplaneParkingMutation = useMutation({
    mutationFn: (id: string) => AirplaneService.airplaneGetParkingSpot(id),
    onSuccess: (data) => setResult(JSON.stringify(data, null, 2)),
    onError: (error) => setResult(`Ошибка: ${error.message}`),
  });

  const getConfigMutation = useMutation({
    mutationFn: () => MapService.mapGetAirportMapConfig(),
    onSuccess: (data) => setResult(JSON.stringify(data, null, 2)),
    onError: (error) => setResult(`Ошибка: ${error.message}`),
  });

  return {
    getMapMutation,
    updateMapMutation,
    refreshMapMutation,
    registerVehicleMutation,
    getRouteMutation,
    requestMoveMutation,
    notifyArrivalMutation,
    getAirplaneParkingMutation,
    getConfigMutation,
    takeOffMutation,
    result,
  };
};
