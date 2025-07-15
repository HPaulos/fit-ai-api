package services

// WorkoutPlanPrompt is the system prompt for generating workout plans
const WorkoutPlanPrompt = `You are an expert fitness trainer with 20+ years experience. Generate comprehensive, personalized workout plans in JSON format only. Return valid JSON matching the exact structure requested.

CORE PRINCIPLES:
- Evidence-based exercise selection
- Progressive overload and safety first
- Balanced muscle groups and movement patterns
- Realistic weights based on fitness level
- Mix of compound and isolation exercises
- Consider user's equipment and goals`

// WorkoutPlanTemplate is the template for generating workout plans
const WorkoutPlanTemplate = `Generate a personalized workout plan for this user:
PROFILE: %s, %s, %s, Fitness: %s, Goals: %v, Equipment: %v, Units: %s
STATS: %d workouts, %d streak, %d total time, %d volume

REQUIREMENTS:
- Create 3-6 workout sessions per week
- Each session: 4-8 exercises with warmups
- Balance push/pull movements across the week
- Include weight, bodyweight, cardio, and flexibility exercises
- Start with compound movements, then isolation
- Use available equipment: %v
- Fitness level: %s

REPS/SETS BY LEVEL:
- Beginner: 3 sets x 12-15 reps (focus on form)
- Intermediate: 4 sets x 8-12 reps (hypertrophy) or 6-8 (strength)
- Advanced: 4-5 sets x 6-8 reps (strength) or 8-12 (hypertrophy)

REST PERIODS:
- Compound: 2-4 minutes (strength) or 1-2 minutes (hypertrophy)
- Isolation: 60-90 seconds

Generate this JSON format:

{
  "id": 1,
  "name": "Professional Plan Name",
  "description": "Comprehensive description of plan approach, methodology, expected results, and timeline.",
  "createdAt": "2024-01-15T00:00:00.000Z",
  "aiFeedbackCycle": 12,
  "planValidityPeriod": 28,
  "sessionsCompleted": 0,
  "planStartDate": "2024-01-15T00:00:00.000Z",
  "hasNewPlanSuggestion": false,
  "suggestedPlan": null,
  "sessions": [
    {
      "id": "session_1",
      "name": "Upper Body Power",
      "note": "Focus on chest, shoulders, triceps. Start compound, then isolation. Rest 2-3 min compound, 60-90 sec isolation.",
      "warmups": [
        {
          "id": "warmup_1",
          "name": "Arm Circles",
          "sets": 2,
          "reps": 10,
          "duration": 30,
          "weight": {"value": 0, "unit": "BODYWEIGHT"},
          "type": "bodyweight",
          "note": "Forward/backward circles to warm up shoulders",
          "equipment": "bodyweight"
        }
      ],
      "exercises": [
        {
          "id": 1,
          "name": "Barbell Bench Press",
          "sets": 4,
          "reps": 8,
          "weight": {"value": 185, "unit": "LB"},
          "type": "weight",
          "equipment": "barbell",
          "note": "Compound chest exercise. Controlled descent, explosive press. Keep feet flat, maintain arch."
        }
      ]
    },
    {
      "id": "session_2",
      "name": "Lower Body Strength",
      "note": "Focus on legs and core. Start squats, then deadlifts. Rest 3-4 min main lifts. Focus form and depth.",
      "warmups": [
        {
          "id": "warmup_2",
          "name": "Bodyweight Squats",
          "sets": 2,
          "reps": 12,
          "weight": {"value": 0, "unit": "BODYWEIGHT"},
          "type": "bodyweight",
          "note": "Focus form and depth to warm up legs",
          "equipment": "bodyweight"
        }
      ],
      "exercises": [
        {
          "id": 2,
          "name": "Barbell Squat",
          "sets": 4,
          "reps": 8,
          "weight": {"value": 225, "unit": "LB"},
          "type": "weight",
          "equipment": "barbell",
          "note": "Keep chest up, knees in line with toes. Go parallel or below."
        }
      ]
    },
    {
      "id": "session_3",
      "name": "Pull Day",
      "note": "Focus back and biceps. Start pull-ups, then rows. Control negative portion. Squeeze shoulder blades.",
      "warmups": [
        {
          "id": "warmup_3",
          "name": "Band Pull-Aparts",
          "sets": 2,
          "reps": 12,
          "weight": {"value": 0, "unit": "BODYWEIGHT"},
          "type": "bodyweight",
          "note": "Shoulder blade activation for pulling movements",
          "equipment": "resistance band"
        }
      ],
      "exercises": [
        {
          "id": 3,
          "name": "Pull-ups",
          "sets": 4,
          "reps": 8,
          "weight": {"value": 0, "unit": "BODYWEIGHT"},
          "type": "bodyweight",
          "equipment": "pull-up bar",
          "note": "Pull chest to bar, control descent. Full range of motion."
        }
      ]
    }
  ]
}

IMPORTANT: Create 3-6 sessions (not 1-2). Each session complete with warmups and exercises. Use specific exercise names with equipment. Include detailed form cues and safety notes. Return only JSON.`
