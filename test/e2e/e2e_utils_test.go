package e2e_test

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/playwright-community/playwright-go"

	"github.com/cooperlutz/go-full/app/config"
)

func newBrowserContextAndPage(t *testing.T, options playwright.BrowserNewContextOptions) (playwright.BrowserContext, playwright.Page) {
	t.Helper()
	context, err := browser.NewContext(options)
	if err != nil {
		t.Fatalf("could not create new context: %v", err)
	}
	t.Cleanup(func() {
		if err := context.Close(); err != nil {
			t.Errorf("could not close context: %v", err)
		}
	})
	p, err := context.NewPage()
	if err != nil {
		t.Fatalf("could not create new page: %v", err)
	}
	return context, p
}

func seedTestData() error {
	ctx := context.Background()

	cfg, err := config.LoadConfigFromEnvVars()
	if err != nil {
		return err
	}

	conn, err := pgx.Connect(ctx, cfg.DB.GetDSN())
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	sqlStatement := `
	INSERT INTO pingpong (pingpong_id,ping_or_pong,created_at,updated_at,deleted_at,deleted) VALUES
	 ('bf7d866a-a41e-4a07-bb0b-84b93f6e196b'::uuid,'pong','2025-11-10 13:41:53.655576-06',NULL,NULL,false),
	 ('cb37277d-1dea-4a59-919a-2f9f88ced1e0'::uuid,'pong','2025-11-10 13:41:54.474201-06','2025-11-10 16:52:17.451034-06','2025-11-10 16:52:17.451034-06',true),
	 ('9cd8684f-3c32-4d82-b072-d5c5b2be5d14'::uuid,'pong','2025-11-10 13:42:15.155329-06','2025-11-10 16:52:17.451034-06',NULL,false),
	 ('d26917af-ad61-4222-b951-e1c8085944b0'::uuid,'pong','2025-11-10 13:42:15.973667-06',NULL,NULL,false),
	 ('6e0a4959-5279-4c44-ba02-890be35e6a86'::uuid,'pong','2025-11-10 13:43:00.219618-06',NULL,NULL,false),
	 ('b9a02664-61bf-4f5e-9ea3-f58b931eb305'::uuid,'ping','2025-11-10 13:43:01.074746-06','2025-11-10 16:52:17.451034-06',NULL,false),
	 ('423b559c-edc4-4070-9f55-07322a3e5cf0'::uuid,'ping','2025-11-10 13:49:19.385166-06','2025-11-10 16:52:17.451034-06','2025-11-10 16:52:17.451034-06',true),
	 ('5f41e894-b63c-48e3-a4e2-07fb3df0068b'::uuid,'ping','2025-11-10 13:49:20.123703-06',NULL,NULL,false),
	 ('2e8993b7-56c9-4363-a291-bc3e86c108a5'::uuid,'ping','2025-11-10 16:52:16.798368-06',NULL,NULL,false),
	 ('d1d22514-5f53-4d40-890b-5350602b3b5f'::uuid,'ping','2025-11-10 16:52:17.451034-06',NULL,NULL,false)
     ON CONFLICT (pingpong_id) DO NOTHING;
	 
	INSERT INTO exam_library.exams
	(exam_id, created_at, updated_at, deleted_at, deleted, "name", grade_level)
	VALUES('11111111-1111-1111-1111-111111111111'::uuid, '2025-11-10 16:52:17.451034-06', NULL, NULL, false, 'Sample Exam', 5)
	ON CONFLICT (exam_id) DO NOTHING;

	INSERT INTO exam_library.exam_questions
	(exam_question_id, created_at, updated_at, deleted_at, deleted, exam_id, "index", question_text, answer_text, question_type, possible_points, response_options)
	VALUES('22222222-2222-2222-2222-222222222222'::uuid, '2025-11-10 16:52:17.451034-06', NULL, NULL, false, '11111111-1111-1111-1111-111111111111'::uuid, 
	 1, 'What is the capital of France?', 'Paris', 'multiple-choice', 5,  ARRAY['Berlin', 'Madrid', 'Paris', 'Rome']::text[])
	 ON CONFLICT (exam_question_id) DO NOTHING;
	 `

	_, err = conn.Exec(ctx, sqlStatement)
	if err != nil {
		return err
	}
	return nil
}
