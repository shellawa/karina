import type { common, models } from "./wailsjs/go/models"

export type ParticipantSolveDetails = {
  acCount: number
  totalTime: number
  totalMemory: number
  minTime: number
  maxTime: number
  minMemory: number
  maxMemory: number
  status: string
  res: common.TestSolveResult[]
  id: string
  name: string
  organization: string
  code?: string
  testCases: models.TestCase[]
}
