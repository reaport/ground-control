import { useQuery } from "@tanstack/react-query";
import { MapService } from "../../api";
import AirportGraph from "../AirportGraph/AirportGraph";

const AdminPanel = () => {
  const { data: map, refetch: refetchMap } = useQuery({
    queryKey: ["adminPanel"],
    queryFn: MapService.mapGetAirportMap,
  });

  return (
    <pre>
      <button onClick={() => refetchMap()}>Refetch</button>
      {map! && <AirportGraph data={map} />}
    </pre>
  );
};

export default AdminPanel;
