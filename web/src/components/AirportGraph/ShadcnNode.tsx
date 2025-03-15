import React from "react"
import { Handle, Position } from "reactflow"

// Example ShadCN Card imports (adjust paths to your setup)
import {
  Card,
  CardHeader,
  CardTitle,
  CardDescription,
} from "@/components/ui/card"

// VehicleType and Vehicle are your own types, but for the sake
// of demonstration we assume these are from your openapi-typescript-codegen
import type { Vehicle, VehicleType } from "../../api" 

type ShadcnNodeProps = {
  data: {
    label: string
    types: VehicleType[]
    vehicles?: Vehicle[]
  }
}

export function ShadcnNode({ data }: ShadcnNodeProps) {
  return (
    <Card className="w-[200px]">
      <CardHeader>
        <CardTitle>{data.label}</CardTitle>
        {data.types?.length ? (
          <CardDescription>{data.types.join(", ")}</CardDescription>
        ) : null}
        {/* If you want to list vehicles: 
          data.vehicles?.length ? (
            <CardDescription>
              Vehicles: {data.vehicles.map(v => v.id).join(", ")}
            </CardDescription>
          ) : null
        */}
      </CardHeader>

      {/* React Flow "handles" for connecting nodes */}
      <Handle type="target" position={Position.Top} />
      <Handle type="source" position={Position.Bottom} />
    </Card>
  )
}
