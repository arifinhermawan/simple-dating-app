CREATE TABLE IF NOT EXISTS public."user" (
	id bigint GENERATED ALWAYS AS IDENTITY NOT NULL,
	username varchar(20) NOT NULL,
	"password" varchar NOT NULL,
	CONSTRAINT user_pk PRIMARY KEY (id),
	CONSTRAINT user_unique UNIQUE (username)
);

CREATE TABLE IF NOT EXISTS public.profile (
	user_id bigint NOT NULL,
	username varchar(20) NOT NULL,
	is_verified boolean DEFAULT false NOT NULL,
	is_infinite_scroll boolean DEFAULT false NOT NULL,
	swipe_count int DEFAULT 0 NULL,
	last_swipe timestamp NULL,
    CONSTRAINT profile_unique UNIQUE (user_id),
	CONSTRAINT profile_user_fk FOREIGN KEY (user_id) REFERENCES public."user"(id) ON DELETE CASCADE,
	CONSTRAINT profile_user_fk_1 FOREIGN KEY (username) REFERENCES public."user"(username) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS public.swipe_history (
	id bigint GENERATED ALWAYS AS IDENTITY NOT NULL,
	swiper_id bigint NOT NULL,
	swiped_id bigint NOT NULL,
	swipe_direction char(5) NOT NULL,
	created_at timestamp DEFAULT now() NOT NULL,
	CONSTRAINT swipe_history_pk PRIMARY KEY (id),
	CONSTRAINT swipe_history_profile_fk FOREIGN KEY (swiper_id) REFERENCES public.profile(user_id),
	CONSTRAINT swipe_history_profile_fk_1 FOREIGN KEY (swiped_id) REFERENCES public.profile(user_id)
);

CREATE TABLE IF NOT EXISTS public.purchase_history (
	id bigint GENERATED ALWAYS AS IDENTITY NOT NULL,
	user_id bigint NOT NULL,
	upgrade_type smallint NOT NULL,
	created_at timestamp DEFAULT now() NOT NULL,
	CONSTRAINT purchase_history_pk PRIMARY KEY (id),
	CONSTRAINT purchase_history_profile_fk FOREIGN KEY (user_id) REFERENCES public.profile(user_id)
);

