export type LiveSessionStatus = 'active' | 'finalizing' | 'completed' | 'cancelled';

export interface LiveSession {
  id: string;
  title?: string;
  status: LiveSessionStatus;
  chunk_count: number;
  last_sequence: number;
  accumulated_transcript?: string;
  output_audio_path?: string;
  final_job_id?: string;
  created_at: string;
  updated_at: string;
}

export interface LiveChunkSegment {
  start: number;
  end: number;
  text: string;
  speaker?: string;
}

export interface LiveChunk {
  sequence: number;
  start_offset: number;
  end_offset: number;
  text: string;
  segments?: LiveChunkSegment[];
}

export interface LiveStreamEvent {
  type: 'snapshot' | 'chunk' | 'status';
  session_id: string;
  session_status: LiveSessionStatus;
  title?: string;
  chunk?: LiveChunk;
  chunks?: LiveChunk[];
  accumulated_text?: string;
  final_job_id?: string;
  timestamp: string;
}
