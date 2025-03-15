import React, { useState } from "react";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import {
  Command,
  CommandInput,
  CommandItem,
  CommandList,
} from "@/components/ui/command";
import { AirportMap, VehicleType } from "@/api";
import { useFetchFunctions } from "@/hooks/useFetchFunctions";
import { PacmanLoader } from "react-spinners";
import { Popover, PopoverContent, PopoverTrigger } from "../ui/popover";

interface TransportControlPanelProps {
  airportMap: AirportMap;
}

const TransportControlPanel: React.FC<TransportControlPanelProps> = ({
  airportMap,
}) => {
  const [from, setFrom] = useState("");
  const [to, setTo] = useState("");
  const [vehicleType, setVehicleType] = useState<VehicleType>(
    VehicleType.AIRPLANE
  );
  const [vehicleId, setVehicleId] = useState("");
  const [nodeId, setNodeId] = useState("");

  const {
    getMapMutation,
    refreshMapMutation,
    registerVehicleMutation,
    getRouteMutation,
    requestMoveMutation,
    notifyArrivalMutation,
    getAirplaneParkingMutation,
    getConfigMutation,
    takeOffMutation,
    result,
  } = useFetchFunctions();

  const isGeneralPending =
    getMapMutation.isPending ||
    refreshMapMutation.isPending ||
    registerVehicleMutation.isPending ||
    getRouteMutation.isPending ||
    requestMoveMutation.isPending ||
    notifyArrivalMutation.isPending ||
    getAirplaneParkingMutation.isPending ||
    getConfigMutation.isPending ||
    takeOffMutation.isPending;

  return (
    <div className="p-6 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
      <Card className="flex flex-col h-full">
        <CardHeader>
          <CardTitle>Запрос карты</CardTitle>
        </CardHeader>
        <CardContent className="flex-grow"></CardContent>
        <CardFooter className="mt-auto">
          <Button onClick={() => getMapMutation.mutate()}>
            Запросить карту
          </Button>
        </CardFooter>
      </Card>
      <Card className="flex flex-col h-full">
        <CardHeader>
          <CardTitle>Очистить карту</CardTitle>
        </CardHeader>
        <CardContent className="flex-grow"></CardContent>
        <CardFooter className="mt-auto">
          <Button onClick={() => refreshMapMutation.mutate()}>
            Очистить карту
          </Button>
        </CardFooter>
      </Card>
      <Card>
        <CardHeader>
          <CardTitle>Регистрация транспорта</CardTitle>
        </CardHeader>
        <CardContent>
          <Select
            value={vehicleType}
            onValueChange={(value) => setVehicleType(value as VehicleType)}
          >
            <SelectTrigger>
              <SelectValue placeholder="Выберите тип транспорта" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value={VehicleType.AIRPLANE}>Самолет</SelectItem>
              <SelectItem value={VehicleType.CATERING}>Кейтеринг</SelectItem>
              <SelectItem value={VehicleType.REFUELING}>Заправка</SelectItem>
              <SelectItem value={VehicleType.CLEANING}>Уборка</SelectItem>
              <SelectItem value={VehicleType.BAGGAGE}>Багаж</SelectItem>
              <SelectItem value={VehicleType.FOLLOW_ME}>Follow-me</SelectItem>
              <SelectItem value={VehicleType.CHARGING}>Зарядка</SelectItem>
              <SelectItem value={VehicleType.BUS}>Автобус</SelectItem>
              <SelectItem value={VehicleType.RAMP}>Платформа</SelectItem>
            </SelectContent>
          </Select>
        </CardContent>
        <CardFooter>
          <Button onClick={() => registerVehicleMutation.mutate(vehicleType)}>
            Зарегистрировать транспорт
          </Button>
        </CardFooter>
      </Card>
      <Card>
        <CardHeader>
          <CardTitle>Запрос маршрута</CardTitle>
        </CardHeader>
        <CardContent className="flex flex-col gap-2">
          <Popover>
            <PopoverTrigger asChild>
              <Button variant="outline" className="w-full text-left">
                {from || "Откуда"}
              </Button>
            </PopoverTrigger>
            <PopoverContent className="w-full p-0">
              <Command>
                <CommandInput placeholder="Поиск..." />
                <CommandList>
                  {airportMap.nodes.map((node) => (
                    <CommandItem
                      key={node.id}
                      value={node.id}
                      onSelect={() => setFrom(node.id)}
                    >
                      {node.id}
                    </CommandItem>
                  ))}
                </CommandList>
              </Command>
            </PopoverContent>
          </Popover>
          <Popover>
            <PopoverTrigger asChild>
              <Button variant="outline" className="w-full text-left">
                {to || "Куда"}
              </Button>
            </PopoverTrigger>
            <PopoverContent className="w-full p-0">
              <Command>
                <CommandInput placeholder="Поиск..." />
                <CommandList>
                  {airportMap.nodes.map((node) => (
                    <CommandItem
                      key={node.id}
                      value={node.id}
                      onSelect={() => setTo(node.id)}
                    >
                      {node.id}
                    </CommandItem>
                  ))}
                </CommandList>
              </Command>
            </PopoverContent>
          </Popover>
        </CardContent>
        <CardFooter>
          <Button
            onClick={() =>
              getRouteMutation.mutate({ from, to, type: vehicleType })
            }
          >
            Запросить маршрут
          </Button>
        </CardFooter>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle>Разрешение на перемещение</CardTitle>
        </CardHeader>
        <CardContent className="flex flex-col gap-2">
          <Input
            placeholder="ID транспорта"
            value={vehicleId}
            onChange={(e) => setVehicleId(e.target.value)}
          />
          <Popover>
            <PopoverTrigger asChild>
              <Button variant="outline" className="w-full text-left">
                {from || "Откуда"}
              </Button>
            </PopoverTrigger>
            <PopoverContent className="w-full p-0">
              <Command>
                <CommandInput placeholder="Поиск..." />
                <CommandList>
                  {airportMap.nodes.map((node) => (
                    <CommandItem
                      key={node.id}
                      value={node.id}
                      onSelect={() => setFrom(node.id)}
                    >
                      {node.id}
                    </CommandItem>
                  ))}
                </CommandList>
              </Command>
            </PopoverContent>
          </Popover>
          <Popover>
            <PopoverTrigger asChild>
              <Button variant="outline" className="w-full text-left">
                {to || "Куда"}
              </Button>
            </PopoverTrigger>
            <PopoverContent className="w-full p-0">
              <Command>
                <CommandInput placeholder="Поиск..." />
                <CommandList>
                  {airportMap.nodes.map((node) => (
                    <CommandItem
                      key={node.id}
                      value={node.id}
                      onSelect={() => setTo(node.id)}
                    >
                      {node.id}
                    </CommandItem>
                  ))}
                </CommandList>
              </Command>
            </PopoverContent>
          </Popover>
        </CardContent>
        <CardFooter>
          <Button
            onClick={() =>
              requestMoveMutation.mutate({ vehicleId, vehicleType, from, to })
            }
          >
            Запросить перемещение
          </Button>
        </CardFooter>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle>Уведомление о прибытии</CardTitle>
        </CardHeader>
        <CardContent className="flex flex-col gap-2">
          <Input
            placeholder="ID транспорта"
            value={vehicleId}
            onChange={(e) => setVehicleId(e.target.value)}
          />
          <Popover>
            <PopoverTrigger asChild>
              <Button variant="outline" className="w-full text-left">
                {nodeId || "ID узла"}
              </Button>
            </PopoverTrigger>
            <PopoverContent className="w-full p-0">
              <Command>
                <CommandInput placeholder="Поиск..." />
                <CommandList>
                  {airportMap.nodes.map((node) => (
                    <CommandItem
                      key={node.id}
                      value={node.id}
                      onSelect={() => setNodeId(node.id)}
                    >
                      {node.id}
                    </CommandItem>
                  ))}
                </CommandList>
              </Command>
            </PopoverContent>
          </Popover>
        </CardContent>
        <CardFooter>
          <Button
            onClick={() =>
              notifyArrivalMutation.mutate({ vehicleId, vehicleType, nodeId })
            }
          >
            Уведомить о прибытии
          </Button>
        </CardFooter>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle>Парковочное место</CardTitle>
        </CardHeader>
        <CardContent>
          <Input
            placeholder="ID транспорта"
            value={vehicleId}
            onChange={(e) => setVehicleId(e.target.value)}
          />
        </CardContent>
        <CardFooter>
          <Button onClick={() => getAirplaneParkingMutation.mutate(vehicleId)}>
            Найти парковку
          </Button>
        </CardFooter>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle>Получение конфигурации</CardTitle>
        </CardHeader>
        <CardContent></CardContent>
        <CardFooter>
          <Button onClick={() => getConfigMutation.mutate()}>Получить</Button>
        </CardFooter>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle>Вылет самолета</CardTitle>
        </CardHeader>
        <CardContent>
          <Input
            placeholder="ID транспорта"
            value={vehicleId}
            onChange={(e) => setVehicleId(e.target.value)}
          />
        </CardContent>
        <CardFooter>
          <Button onClick={() => takeOffMutation.mutate(vehicleId)}>
            Вылететь
          </Button>
        </CardFooter>
      </Card>
      {/* Вывод результата запроса */}
      <Card className="col-span-full">
        <CardHeader>
          <CardTitle>Результат</CardTitle>
        </CardHeader>
        <CardContent>
          {isGeneralPending ? (
            <PacmanLoader />
          ) : (
            <pre className="whitespace-pre-wrap break-words text-sm p-2 bg-gray-100 rounded">
              {!result ? "УСПЕХ" : result}
            </pre>
          )}
        </CardContent>
      </Card>
    </div>
  );
};

export default TransportControlPanel;
